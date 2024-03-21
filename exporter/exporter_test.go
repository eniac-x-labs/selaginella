package exporter

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evm-layer2/selaginella/bindings"
	"math/big"
	"strings"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	client, _ := ethclient.DialContext(context.Background(), "https://eth-sepolia.g.alchemy.com/v2/13PoJNZA67XxE87RJ8ZFyGkWKRGcOOgY")
	latestBlockNumber, _ := client.BlockNumber(context.Background())

	hex := "0f0b59bddf091da85ebee2d547b8e8c2d2a92fa23982bc54fe13d6e439b5f4e8"
	hex = strings.TrimPrefix(hex, "0x")
	key, _ := crypto.HexToECDSA(hex)
	priKey := key

	cOpts := &bind.CallOpts{
		BlockNumber: new(big.Int).SetUint64(latestBlockNumber),
		From:        crypto.PubkeyToAddress(priKey.PublicKey),
	}

	//var newPools []bindings.IL1PoolManagerPool
	//var newPool bindings.IL1PoolManagerPool
	//newPool.Token = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	//newPool.TotalAmount = new(big.Int).SetUint64(0)
	//newPool.TotalFeeClaimed = new(big.Int).SetUint64(0)
	//newPool.TotalFee = new(big.Int).SetUint64(0)
	//
	//newPools = append(newPools, newPool)
	//fmt.Println(newPools)
	//topts, err := bind.NewKeyedTransactorWithChainID(priKey, new(big.Int).SetUint64(11155111))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//topts.Context = context.Background()
	//topts.NoSend = true
	//
	l1PoolContract, _ := bindings.NewL1PoolManager(common.HexToAddress("0x4F34C922fB0D80c7d79Ac25e497d90d7efa513C2"), client)
	//tx, _ := l1PoolContract.CompletePoolAndNew(topts, newPools)
	//
	//l1Parsed, err := abi.JSON(strings.NewReader(
	//	bindings.L1PoolManagerABI,
	//))
	//rawL1PoolContract := bind.NewBoundContract(
	//	common.HexToAddress("0x1DE4c1C613aA0Ba3F52eEa56D2D0632e252B9E5F"), l1Parsed, client, client,
	//	client,
	//)
	//fmt.Println(tx)
	//finalTx, err := rawL1PoolContract.RawTransact(topts, tx.Data())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = client.SendTransaction(context.Background(), finalTx)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(finalTx.Hash())

	length, _ := l1PoolContract.GetPoolLength(cOpts, common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"))
	fmt.Println(length.String())
	balance, _ := l1PoolContract.GetPool(cOpts, common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"), new(big.Int).Sub(length, new(big.Int).SetUint64(2)))
	fmt.Println(int64(balance.EndTimestamp) > time.Now().Unix())
	fmt.Println(balance)
}

func Test3(t *testing.T) {
	fmt.Println(common.HexToHash("0x0000000000000000000000000000000000000000000000004563918244f40000").Big().String())
}
