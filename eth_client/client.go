package node

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/evm-layer2/selaginella/eth_client/retry"
)

const (
	defaultDialTimeout    = 5 * time.Second
	defaultDialAttempts   = 5
	defaultRequestTimeout = 10 * time.Second
)

type EthClient interface {
	TxReceiptByHash(common.Hash) (*types.Transaction, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	Close()
}

type clnt struct {
	rpc RPC
}

func DialEthClient(ctx context.Context, rpcUrl string) (EthClient, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()

	bOff := retry.Exponential()
	rpcClient, err := retry.Do(ctx, defaultDialAttempts, bOff, func() (*rpc.Client, error) {
		if !IsURLAvailable(rpcUrl) {
			return nil, fmt.Errorf("address unavailable (%s)", rpcUrl)
		}

		client, err := rpc.DialContext(ctx, rpcUrl)
		if err != nil {
			return nil, fmt.Errorf("failed to dial address (%s): %w", rpcUrl, err)
		}

		return client, nil
	})

	if err != nil {
		return nil, err
	}

	return &clnt{rpc: NewRPC(rpcClient)}, nil
}

func (c *clnt) TxReceiptByHash(hash common.Hash) (*types.Transaction, error) {
	ctxwt, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var tx *types.Transaction
	err := c.rpc.CallContext(ctxwt, &tx, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, err
	} else if tx == nil {
		return nil, ethereum.NotFound
	}

	return tx, nil
}

// SendTransaction injects a signed transaction into the pending pool for execution.
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (c *clnt) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	data, err := tx.MarshalBinary()
	if err != nil {
		return err
	}
	return c.rpc.CallContext(ctx, nil, "eth_sendRawTransaction", hexutil.Encode(data))
}

func (c *clnt) Close() {
	c.rpc.Close()
}

type RPC interface {
	Close()
	CallContext(ctx context.Context, result any, method string, args ...any) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
}

type rpcClient struct {
	rpc *rpc.Client
}

func NewRPC(client *rpc.Client) RPC {
	return &rpcClient{client}
}

func (c *rpcClient) Close() {
	c.rpc.Close()
}

func (c *rpcClient) CallContext(ctx context.Context, result any, method string, args ...any) error {
	err := c.rpc.CallContext(ctx, result, method, args...)
	return err
}

func (c *rpcClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	err := c.rpc.BatchCallContext(ctx, b)
	return err
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	return rpc.BlockNumber(number.Int64()).String()
}

func toFilterArg(q ethereum.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{"address": q.Addresses, "topics": q.Topics}
	if q.BlockHash != nil {
		arg["blockHash"] = *q.BlockHash
		if q.FromBlock != nil || q.ToBlock != nil {
			return nil, errors.New("cannot specify both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if q.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = toBlockNumArg(q.FromBlock)
		}
		arg["toBlock"] = toBlockNumArg(q.ToBlock)
	}
	return arg, nil
}

func IsURLAvailable(address string) bool {
	u, err := url.Parse(address)
	if err != nil {
		return false
	}
	addr := u.Host
	if u.Port() == "" {
		switch u.Scheme {
		case "http", "ws":
			addr += ":80"
		case "https", "wss":
			addr += ":443"
		default:
			return true
		}
	}
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func DialEthClientWithTimeout(ctx context.Context, url string, disableHTTP2 bool) (
	*ethclient.Client, error) {
	ctxt, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()
	if strings.HasPrefix(url, "http") {
		httpClient := new(http.Client)
		if disableHTTP2 {
			log.Debug("Disabled HTTP/2 support in  eth client")
			httpClient.Transport = &http.Transport{
				TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
			}
		}
		rpcClient, err := rpc.DialHTTPWithClient(url, httpClient)
		if err != nil {
			return nil, err
		}
		return ethclient.NewClient(rpcClient), nil
	}
	return ethclient.DialContext(ctxt, url)
}
