package exporter

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	kms "cloud.google.com/go/kms/apiv1"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/evm-layer2/selaginella/bindings"
	common2 "github.com/evm-layer2/selaginella/common"
	"github.com/evm-layer2/selaginella/config"
	"github.com/evm-layer2/selaginella/database"
	node "github.com/evm-layer2/selaginella/eth_client"
	"github.com/evm-layer2/selaginella/sign"
)

type Exporter struct {
	exporterConfig *config.Exporter
	db             *database.DB
	shutdown       context.CancelCauseFunc
	ethClients     map[uint64]node.EthClient
	*HsmConfig
	poolAddresses       map[uint64]common.Address
	EthAddress          map[uint64]common.Address
	WEthAddress         map[uint64]common.Address
	USDTAddress         map[uint64]common.Address
	RawL1BridgeContract *bind.BoundContract
	RawL2BridgeContract map[uint64]*bind.BoundContract
	L1BridgeContract    *bindings.L1PoolManager
	L2BridgeContract    map[uint64]*bindings.L2PoolManager
	privateKey          *ecdsa.PrivateKey
	l1ChainID           uint64
	stopped             atomic.Bool
	*TransferMultiple
}

type HsmConfig struct {
	EnableHsm  bool
	HsmAPIName string
	HsmCreden  string
	HsmAddress string
}

type TransferMultiple struct {
	L1Multiple uint64
	L2Multiple uint64
}

func NewExporter(ctx context.Context, exporterConfig config.Exporter, hsmCfg *HsmConfig, db *database.DB, chainRpcCfg []*config.RPC, shutdown context.CancelCauseFunc, priKey *ecdsa.PrivateKey, l1ChainID uint64, multipleCfg *TransferMultiple) (*Exporter, error) {
	ethClients := make(map[uint64]node.EthClient)
	var rawL1BridgeContract *bind.BoundContract
	rawL2BridgeContracts := make(map[uint64]*bind.BoundContract)
	var l1BridgeContract *bindings.L1PoolManager
	l2BridgeContracts := make(map[uint64]*bindings.L2PoolManager)
	poolAddresses := make(map[uint64]common.Address)
	EthAddress := make(map[uint64]common.Address)
	WEthAddress := make(map[uint64]common.Address)
	USDTAddress := make(map[uint64]common.Address)

	for i := range chainRpcCfg {
		if chainRpcCfg[i].ChainId == l1ChainID {
			client, err := node.DialEthClient(ctx, chainRpcCfg[i].RpcUrl)
			if err != nil {
				log.Errorln("dial client fail", "error", err.Error())
				return nil, err
			}
			ethClients[chainRpcCfg[i].ChainId] = client

			l1Parsed, err := abi.JSON(strings.NewReader(
				bindings.L1PoolManagerABI,
			))
			if err != nil {
				log.Errorln("selaginella parse l1 pool contract abi fail", "err", err)
				return nil, err
			}

			l1Client, _ := node.DialEthClientWithTimeout(ctx, chainRpcCfg[i].RpcUrl, false)
			rawL1PoolContract := bind.NewBoundContract(
				common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress), l1Parsed, l1Client, l1Client,
				l1Client,
			)
			rawL1BridgeContract = rawL1PoolContract

			l1PoolContract, err := bindings.NewL1PoolManager(common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress), l1Client)
			l1BridgeContract = l1PoolContract
			poolAddresses[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress)
			EthAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].EthAddress)
			WEthAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].WEthAddress)
			USDTAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].USDTAddress)

		} else {
			client, err := node.DialEthClient(ctx, chainRpcCfg[i].RpcUrl)
			if err != nil {
				log.Errorln("dial client fail", "error", err.Error())
				return nil, err
			}
			ethClients[chainRpcCfg[i].ChainId] = client

			l2Parsed, err := abi.JSON(strings.NewReader(
				bindings.L2PoolManagerABI,
			))
			if err != nil {
				log.Errorln("selaginella parse l2 pool contract abi fail", "err", err)
				return nil, err
			}

			l2Client, _ := node.DialEthClientWithTimeout(ctx, chainRpcCfg[i].RpcUrl, false)
			rawL2PoolContract := bind.NewBoundContract(
				common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress), l2Parsed, l2Client, l2Client,
				l2Client,
			)
			rawL2BridgeContracts[chainRpcCfg[i].ChainId] = rawL2PoolContract

			l2PoolContract, err := bindings.NewL2PoolManager(common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress), l2Client)
			l2BridgeContracts[chainRpcCfg[i].ChainId] = l2PoolContract
			poolAddresses[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].FoundingPoolAddress)
			EthAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].EthAddress)
			WEthAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].WEthAddress)
			USDTAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].USDTAddress)
		}

	}

	return &Exporter{
		exporterConfig:      &exporterConfig,
		db:                  db,
		shutdown:            shutdown,
		ethClients:          ethClients,
		HsmConfig:           hsmCfg,
		RawL1BridgeContract: rawL1BridgeContract,
		RawL2BridgeContract: rawL2BridgeContracts,
		L1BridgeContract:    l1BridgeContract,
		L2BridgeContract:    l2BridgeContracts,
		poolAddresses:       poolAddresses,
		privateKey:          priKey,
		EthAddress:          EthAddress,
		WEthAddress:         WEthAddress,
		USDTAddress:         USDTAddress,
		l1ChainID:           l1ChainID,
		TransferMultiple:    multipleCfg,
	}, nil
}

func (er *Exporter) Start(ctx context.Context) error {

	log.Infoln("exporter config", er.exporterConfig.ExportAddress)
	log.Infoln("Starting selaginella_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<head><title>Selaginella Exporter</title></head>
		<body>
		<h1>Selaginella Exporter</h1>
		<p><a href="/metrics">Metrics</a></p>
		</body>
		</html>`))
	})

	go er.metricEthFundBalance()

	go er.metricWEthFundBalance()

	go er.metricUSDTFundBalance()

	go er.metricEthereumPoolBalance()

	log.Infoln("Listening on", er.exporterConfig.ExportAddress)
	if err := http.ListenAndServe(er.exporterConfig.ExportAddress, nil); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (er *Exporter) Stop(ctx context.Context) error {
	var result error

	if er.db != nil {
		if err := er.db.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close DB: %w", err))
		}
	}

	er.stopped.Store(true)
	log.Info("exporter service shutdown complete")

	return nil
}

func (er *Exporter) Stopped() bool {
	return er.stopped.Load()
}

func (er *Exporter) metricEthFundBalance() {
	ticker := time.NewTicker(2 * time.Hour)
	var ethereumPoolBalance *big.Int
	var ethereumChainId uint64
	var opPoolBalance *big.Int
	var opChainId uint64
	var polygonZkEvmPoolBalance *big.Int
	var polygonZkEvmChainId uint64
	var scrollPoolBalance *big.Int
	var scrollChainId uint64
	var chainCount uint64

	for {
		<-ticker.C
		for chainId, client := range er.ethClients {
			latestBlockNumber, err := client.GetLatestBlock()
			if err != nil {
				log.Errorln("exporter", "client get latest block number fail", err)
				return
			}

			balance, err := client.GetBalanceByBlockNumber(er.poolAddresses[chainId].String(), latestBlockNumber)
			if err != nil {
				log.Errorln("exporter", "client get pool balance fail", err)
				return
			}
			switch chainId {
			case common2.ChainEthID, common2.ChainEthSepoliaID:
				ethereumChainId = chainId
				ethereumPoolBalance = balance
			case common2.ChainOpID, common2.ChainOpSepoliaID:
				opChainId = chainId
				opPoolBalance = balance
			case common2.ChainPolygonZkEvmID, common2.ChainPolygonZkEvmSepoliaID:
				polygonZkEvmChainId = chainId
				polygonZkEvmPoolBalance = balance
			case common2.ChainScrollID, common2.ChainScrollSepoliaID:
				scrollChainId = chainId
				scrollPoolBalance = balance
			default:
				log.Errorln("unknown chain")
				return
			}
			chainCount++
		}
		ChainTotalBalance := new(big.Int).Add(ethereumPoolBalance, opPoolBalance)
		ChainTotalBalance.Add(ChainTotalBalance, polygonZkEvmPoolBalance)
		ChainTotalBalance.Add(ChainTotalBalance, scrollPoolBalance)
		ChainAverageBalance := new(big.Int).Div(ChainTotalBalance, new(big.Int).SetUint64(chainCount))

		if err := er.metricOpBalance(&opPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, opChainId, ethereumChainId, er.EthAddress[ethereumChainId].String(), er.EthAddress[opChainId].String()); err != nil {
			log.Errorln("metric op eth balance fail", "error", err)
		}
		if err := er.metricPolygonZkEvmBalance(&polygonZkEvmPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, polygonZkEvmChainId, ethereumChainId, er.EthAddress[ethereumChainId].String(), er.EthAddress[polygonZkEvmChainId].String()); err != nil {
			log.Errorln("metric polygonZkEvm eth balance fail", "error", err)
		}
		if err := er.metricScrollBalance(&scrollPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, scrollChainId, ethereumChainId, er.EthAddress[ethereumChainId].String(), er.EthAddress[scrollChainId].String()); err != nil {
			log.Errorln("metric scroll eth balance fail", "error", err)
		}

		ethPBF, _ := ethereumPoolBalance.Float64()
		opPBF, _ := opPoolBalance.Float64()
		pyPBF, _ := polygonZkEvmPoolBalance.Float64()
		scPBF, _ := scrollPoolBalance.Float64()
		averageBalanceF, _ := ChainAverageBalance.Float64()

		ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[ethereumChainId].String(), er.EthAddress[ethereumChainId].String()).Set(ethPBF)
		opPoolBalanceMetric.WithLabelValues(er.poolAddresses[opChainId].String(), er.EthAddress[opChainId].String()).Set(opPBF)
		polygonZkEvmPoolBalanceMetric.WithLabelValues(er.poolAddresses[polygonZkEvmChainId].String(), er.EthAddress[polygonZkEvmChainId].String()).Set(pyPBF)
		scrollPoolBalanceMetric.WithLabelValues(er.poolAddresses[scrollChainId].String(), er.EthAddress[scrollChainId].String()).Set(scPBF)
		chainAverageBalanceMetric.WithLabelValues("ETH").Set(averageBalanceF)
	}
}

func (er *Exporter) metricWEthFundBalance() {
	ticker := time.NewTicker(2 * time.Hour)
	var ethereumPoolBalance *big.Int
	var ethereumChainId uint64
	var opPoolBalance *big.Int
	var opChainId uint64
	var polygonZkEvmPoolBalance *big.Int
	var polygonZkEvmChainId uint64
	var scrollPoolBalance *big.Int
	var scrollChainId uint64
	var chainCount uint64

	for {
		<-ticker.C
		for chainId, client := range er.ethClients {
			latestBlockNumber, err := client.GetLatestBlock()
			if err != nil {
				log.Errorln("exporter", "client get latest block number fail", err)
				return
			}

			balance, err := client.GetERC20Balance(er.WEthAddress[chainId], er.poolAddresses[chainId], latestBlockNumber)
			if err != nil {
				log.Errorln("exporter", "client get pool balance fail", err)
				return
			}
			switch chainId {
			case common2.ChainEthID, common2.ChainEthSepoliaID:
				ethereumChainId = chainId
				ethereumPoolBalance = balance
			case common2.ChainOpID, common2.ChainOpSepoliaID:
				opChainId = chainId
				opPoolBalance = balance
			case common2.ChainPolygonZkEvmID, common2.ChainPolygonZkEvmSepoliaID:
				polygonZkEvmChainId = chainId
				polygonZkEvmPoolBalance = balance
			case common2.ChainScrollID, common2.ChainScrollSepoliaID:
				scrollChainId = chainId
				scrollPoolBalance = balance
			default:
				log.Errorln("unknown chain")
				return
			}
			chainCount++
		}
		ChainTotalBalance := new(big.Int).Add(ethereumPoolBalance, opPoolBalance)
		ChainTotalBalance.Add(ChainTotalBalance, polygonZkEvmPoolBalance)
		ChainTotalBalance.Add(ChainTotalBalance, scrollPoolBalance)
		ChainAverageBalance := new(big.Int).Div(ChainTotalBalance, new(big.Int).SetUint64(chainCount))

		if err := er.metricOpBalance(&opPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, opChainId, ethereumChainId, er.WEthAddress[ethereumChainId].String(), er.WEthAddress[opChainId].String()); err != nil {
			log.Errorln("metric op weth balance fail", "error", err)
		}
		if err := er.metricPolygonZkEvmBalance(&polygonZkEvmPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, polygonZkEvmChainId, ethereumChainId, er.WEthAddress[ethereumChainId].String(), er.WEthAddress[polygonZkEvmChainId].String()); err != nil {
			log.Errorln("metric polygonZkEvm weth balance fail", "error", err)
		}
		if err := er.metricScrollBalance(&scrollPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, scrollChainId, ethereumChainId, er.WEthAddress[ethereumChainId].String(), er.WEthAddress[scrollChainId].String()); err != nil {
			log.Errorln("metric scroll weth balance fail", "error", err)
		}

		ethPBF, _ := ethereumPoolBalance.Float64()
		opPBF, _ := opPoolBalance.Float64()
		pyPBF, _ := polygonZkEvmPoolBalance.Float64()
		scPBF, _ := scrollPoolBalance.Float64()
		averageBalanceF, _ := ChainAverageBalance.Float64()

		ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[ethereumChainId].String(), er.WEthAddress[ethereumChainId].String()).Set(ethPBF)
		opPoolBalanceMetric.WithLabelValues(er.poolAddresses[opChainId].String(), er.WEthAddress[opChainId].String()).Set(opPBF)
		polygonZkEvmPoolBalanceMetric.WithLabelValues(er.poolAddresses[polygonZkEvmChainId].String(), er.WEthAddress[polygonZkEvmChainId].String()).Set(pyPBF)
		scrollPoolBalanceMetric.WithLabelValues(er.poolAddresses[scrollChainId].String(), er.WEthAddress[scrollChainId].String()).Set(scPBF)
		chainAverageBalanceMetric.WithLabelValues("WETH").Set(averageBalanceF)
	}
}

func (er *Exporter) metricUSDTFundBalance() {
	ticker := time.NewTicker(2 * time.Hour)
	var ethereumPoolBalance *big.Int
	var ethereumChainId uint64
	var opPoolBalance *big.Int
	var opChainId uint64
	var polygonZkEvmPoolBalance *big.Int
	var polygonZkEvmChainId uint64
	var scrollPoolBalance *big.Int
	var scrollChainId uint64
	var chainCount uint64

	for {
		<-ticker.C
		for chainId, client := range er.ethClients {
			latestBlockNumber, err := client.GetLatestBlock()
			if err != nil {
				log.Errorln("exporter", "client get latest block number fail", err)
				return
			}

			balance, err := client.GetERC20Balance(er.USDTAddress[chainId], er.poolAddresses[chainId], latestBlockNumber)
			if err != nil {
				log.Errorln("exporter", "client get pool balance fail", err)
				return
			}
			switch chainId {
			case common2.ChainEthID, common2.ChainEthSepoliaID:
				ethereumChainId = chainId
				ethereumPoolBalance = balance
			case common2.ChainOpID, common2.ChainOpSepoliaID:
				opChainId = chainId
				opPoolBalance = balance
			case common2.ChainPolygonZkEvmID, common2.ChainPolygonZkEvmSepoliaID:
				polygonZkEvmChainId = chainId
				polygonZkEvmPoolBalance = balance
			case common2.ChainScrollID, common2.ChainScrollSepoliaID:
				scrollChainId = chainId
				scrollPoolBalance = balance
			default:
				log.Errorln("unknown chain")
				return
			}
			chainCount++
		}
		ChainTotalBalance := new(big.Int).Add(ethereumPoolBalance, opPoolBalance)
		ChainTotalBalance.Add(ChainTotalBalance, polygonZkEvmPoolBalance)
		ChainTotalBalance.Add(ChainTotalBalance, scrollPoolBalance)
		ChainAverageBalance := new(big.Int).Div(ChainTotalBalance, new(big.Int).SetUint64(chainCount))

		if err := er.metricOpBalance(&opPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, opChainId, ethereumChainId, er.USDTAddress[ethereumChainId].String(), er.USDTAddress[opChainId].String()); err != nil {
			log.Errorln("metric op usdt balance fail", "error", err)
		}
		if err := er.metricPolygonZkEvmBalance(&polygonZkEvmPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, polygonZkEvmChainId, ethereumChainId, er.USDTAddress[ethereumChainId].String(), er.USDTAddress[polygonZkEvmChainId].String()); err != nil {
			log.Errorln("metric polygonZkEvm usdt balance fail", "error", err)
		}
		if err := er.metricScrollBalance(&scrollPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, scrollChainId, ethereumChainId, er.USDTAddress[ethereumChainId].String(), er.USDTAddress[scrollChainId].String()); err != nil {
			log.Errorln("metric scroll usdt balance fail", "error", err)
		}

		ethPBF, _ := ethereumPoolBalance.Float64()
		opPBF, _ := opPoolBalance.Float64()
		pyPBF, _ := polygonZkEvmPoolBalance.Float64()
		scPBF, _ := scrollPoolBalance.Float64()
		averageBalanceF, _ := ChainAverageBalance.Float64()

		ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[ethereumChainId].String(), er.USDTAddress[ethereumChainId].String()).Set(ethPBF)
		opPoolBalanceMetric.WithLabelValues(er.poolAddresses[opChainId].String(), er.USDTAddress[opChainId].String()).Set(opPBF)
		polygonZkEvmPoolBalanceMetric.WithLabelValues(er.poolAddresses[polygonZkEvmChainId].String(), er.USDTAddress[polygonZkEvmChainId].String()).Set(pyPBF)
		scrollPoolBalanceMetric.WithLabelValues(er.poolAddresses[scrollChainId].String(), er.USDTAddress[scrollChainId].String()).Set(scPBF)
		chainAverageBalanceMetric.WithLabelValues("USDT").Set(averageBalanceF)
	}
}

func (er *Exporter) metricOpBalance(opPoolBalance **big.Int, ethereumPoolBalance **big.Int, ChainAverageBalance *big.Int, TotalChainNum uint64, opChainId uint64, ethereumChainId uint64, l1TokenAddress string, l2TokenAddress string) error {
	ethereumClient := er.ethClients[ethereumChainId]
	opClient := er.ethClients[opChainId]

	if (*opPoolBalance).Cmp(new(big.Int).Div(ChainAverageBalance, new(big.Int).SetUint64(er.L1Multiple))) < 1 {
		transferAmount := new(big.Int).Div(*ethereumPoolBalance, new(big.Int).Div(new(big.Int).SetUint64(TotalChainNum), new(big.Int).SetUint64(2)))
		opts, err := er.getSignOpts(ethereumChainId)
		if err != nil {
			log.Errorln(err)
			return err
		}

		tx, err := er.L1BridgeContract.TransferAssertToBridge(opts, new(big.Int).SetUint64(opChainId), common.HexToAddress(common2.ETH), er.poolAddresses[opChainId], transferAmount)
		if err != nil {
			log.Errorln("get bridge transaction by abi fail", "error", err)
			return err
		}

		finalTx, err := er.RawL1BridgeContract.RawTransact(opts, tx.Data())
		if err != nil {
			log.Errorln("raw send bridge transaction fail", "error", err)
			return err
		}

		ctxwt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = ethereumClient.SendTransaction(ctxwt, finalTx)
		if err != nil {
			log.Errorln("send bridge transaction fail", "error", err)
			return err
		}

		receipt, err := ethereumClient.TxReceiptDetailByHash(finalTx.Hash())
		if err != nil {
			log.Errorln("get transaction by hash fail", "error", err)
			return err
		}

		log.Infoln("send l1 pool eth to op pool transaction success", "txHash", receipt.TxHash)

	}

	if (*opPoolBalance).Cmp(new(big.Int).Mul(ChainAverageBalance, new(big.Int).SetUint64(er.L2Multiple))) > -1 {
		withAmount := new(big.Int).Div(*opPoolBalance, new(big.Int).SetUint64(3))
		opts, err := er.getSignOpts(opChainId)
		if err != nil {
			log.Errorln(err)
			return err
		}

		tx, err := er.L2BridgeContract[opChainId].WithdrawETHtoL1(opts, er.poolAddresses[ethereumChainId], withAmount)
		if err != nil {
			log.Errorln("get bridge transaction by abi fail", "error", err)
			return err
		}

		finalTx, err := er.RawL2BridgeContract[opChainId].RawTransact(opts, tx.Data())
		if err != nil {
			log.Errorln("raw send bridge transaction fail", "error", err)
			return err
		}

		ctxwt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = opClient.SendTransaction(ctxwt, finalTx)
		if err != nil {
			log.Errorln("send bridge transaction fail", "error", err)
		}

		receipt, err := opClient.TxReceiptDetailByHash(finalTx.Hash())
		if err != nil {
			log.Errorln("get transaction by hash fail", "error", err)
			return err
		}

		log.Infoln("withdraw op pool eth to l1 pool transaction success", "txHash", receipt.TxHash)
	}

	l1LastestBlockNumber, _ := ethereumClient.GetLatestBlock()
	l2LastestBlockNumber, _ := opClient.GetLatestBlock()
	var l1PoolBalance *big.Int
	var l2PoolBalance *big.Int

	if l1TokenAddress == common2.ETH {
		l1PoolBalance, _ = ethereumClient.GetBalanceByBlockNumber(er.poolAddresses[ethereumChainId].String(), l1LastestBlockNumber)
		l2PoolBalance, _ = opClient.GetBalanceByBlockNumber(er.poolAddresses[opChainId].String(), l2LastestBlockNumber)
	} else {
		l1PoolBalance, _ = ethereumClient.GetERC20Balance(common.HexToAddress(l1TokenAddress), er.poolAddresses[ethereumChainId], l1LastestBlockNumber)
		l2PoolBalance, _ = opClient.GetERC20Balance(common.HexToAddress(l2TokenAddress), er.poolAddresses[opChainId], l1LastestBlockNumber)
	}

	*ethereumPoolBalance = l1PoolBalance
	*opPoolBalance = l2PoolBalance

	return nil
}

func (er *Exporter) metricPolygonZkEvmBalance(pyPoolBalance **big.Int, ethereumPoolBalance **big.Int, ChainAverageBalance *big.Int, TotalChainNum uint64, pyChainId uint64, ethereumChainId uint64, l1TokenAddress string, l2TokenAddress string) error {
	ethereumClient := er.ethClients[ethereumChainId]
	pyClient := er.ethClients[pyChainId]

	if (*pyPoolBalance).Cmp(new(big.Int).Div(ChainAverageBalance, new(big.Int).SetUint64(er.L1Multiple))) < 1 {
		transferAmount := new(big.Int).Div(*ethereumPoolBalance, new(big.Int).Div(new(big.Int).SetUint64(TotalChainNum), new(big.Int).SetUint64(2)))
		opts, err := er.getSignOpts(ethereumChainId)
		if err != nil {
			log.Errorln(err)
			return err
		}

		tx, err := er.L1BridgeContract.TransferAssertToBridge(opts, new(big.Int).SetUint64(pyChainId), common.HexToAddress(common2.ETH), er.poolAddresses[pyChainId], transferAmount)
		if err != nil {
			log.Errorln("get bridge transaction by abi fail", "error", err)
			return err
		}

		finalTx, err := er.RawL1BridgeContract.RawTransact(opts, tx.Data())
		if err != nil {
			log.Errorln("raw send bridge transaction fail", "error", err)
			return err
		}

		ctxwt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = ethereumClient.SendTransaction(ctxwt, finalTx)
		if err != nil {
			log.Errorln("send bridge transaction fail", "error", err)
			return err
		}

		receipt, err := ethereumClient.TxReceiptDetailByHash(finalTx.Hash())
		if err != nil {
			log.Errorln("get transaction by hash fail", "error", err)
			return err
		}

		log.Infoln("send l1 pool eth to polygon zkEvm pool transaction success", "txHash", receipt.TxHash)

	}

	if (*pyPoolBalance).Cmp(new(big.Int).Mul(ChainAverageBalance, new(big.Int).SetUint64(er.L2Multiple))) > -1 {
		withAmount := new(big.Int).Div(*pyPoolBalance, new(big.Int).SetUint64(3))
		opts, err := er.getSignOpts(pyChainId)
		if err != nil {
			log.Errorln(err)
			return err
		}

		tx, err := er.L2BridgeContract[pyChainId].WithdrawETHtoL1(opts, er.poolAddresses[ethereumChainId], withAmount)
		if err != nil {
			log.Errorln("get bridge transaction by abi fail", "error", err)
			return err
		}

		finalTx, err := er.RawL2BridgeContract[pyChainId].RawTransact(opts, tx.Data())
		if err != nil {
			log.Errorln("raw send bridge transaction fail", "error", err)
			return err
		}

		ctxwt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = pyClient.SendTransaction(ctxwt, finalTx)
		if err != nil {
			log.Errorln("send bridge transaction fail", "error", err)
		}

		receipt, err := pyClient.TxReceiptDetailByHash(finalTx.Hash())
		if err != nil {
			log.Errorln("get transaction by hash fail", "error", err)
			return err
		}

		log.Infoln("withdraw polygon zkEvm pool eth to l1 pool transaction success", "txHash", receipt.TxHash)
	}

	l1LastestBlockNumber, _ := ethereumClient.GetLatestBlock()
	l2LastestBlockNumber, _ := pyClient.GetLatestBlock()
	var l1PoolBalance *big.Int
	var l2PoolBalance *big.Int

	if l1TokenAddress == common2.ETH {
		l1PoolBalance, _ = ethereumClient.GetBalanceByBlockNumber(er.poolAddresses[ethereumChainId].String(), l1LastestBlockNumber)
		l2PoolBalance, _ = pyClient.GetBalanceByBlockNumber(er.poolAddresses[pyChainId].String(), l2LastestBlockNumber)
	} else {
		l1PoolBalance, _ = ethereumClient.GetERC20Balance(common.HexToAddress(l1TokenAddress), er.poolAddresses[ethereumChainId], l1LastestBlockNumber)
		l2PoolBalance, _ = pyClient.GetERC20Balance(common.HexToAddress(l2TokenAddress), er.poolAddresses[pyChainId], l1LastestBlockNumber)
	}

	*ethereumPoolBalance = l1PoolBalance
	*pyPoolBalance = l2PoolBalance

	return nil
}

func (er *Exporter) metricScrollBalance(scrollPoolBalance **big.Int, ethereumPoolBalance **big.Int, ChainAverageBalance *big.Int, TotalChainNum uint64, scrollChainId uint64, ethereumChainId uint64, l1TokenAddress string, l2TokenAddress string) error {
	ethereumClient := er.ethClients[ethereumChainId]
	scrollClient := er.ethClients[scrollChainId]

	if (*scrollPoolBalance).Cmp(new(big.Int).Div(ChainAverageBalance, new(big.Int).SetUint64(er.L1Multiple))) < 1 {
		transferAmount := new(big.Int).Div(*ethereumPoolBalance, new(big.Int).Div(new(big.Int).SetUint64(TotalChainNum), new(big.Int).SetUint64(2)))
		opts, err := er.getSignOpts(ethereumChainId)
		if err != nil {
			log.Errorln(err)
			return err
		}

		tx, err := er.L1BridgeContract.TransferAssertToBridge(opts, new(big.Int).SetUint64(scrollChainId), common.HexToAddress(common2.ETH), er.poolAddresses[scrollChainId], transferAmount)
		if err != nil {
			log.Errorln("get bridge transaction by abi fail", "error", err)
			return err
		}

		finalTx, err := er.RawL1BridgeContract.RawTransact(opts, tx.Data())
		if err != nil {
			log.Errorln("raw send bridge transaction fail", "error", err)
			return err
		}

		ctxwt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = ethereumClient.SendTransaction(ctxwt, finalTx)
		if err != nil {
			log.Errorln("send bridge transaction fail", "error", err)
			return err
		}

		receipt, err := ethereumClient.TxReceiptDetailByHash(finalTx.Hash())
		if err != nil {
			log.Errorln("get transaction by hash fail", "error", err)
			return err
		}

		log.Infoln("send l1 pool eth to scroll pool transaction success", "txHash", receipt.TxHash)

	}

	if (*scrollPoolBalance).Cmp(new(big.Int).Mul(ChainAverageBalance, new(big.Int).SetUint64(er.L2Multiple))) > -1 {
		withAmount := new(big.Int).Div(*scrollPoolBalance, new(big.Int).SetUint64(3))
		opts, err := er.getSignOpts(scrollChainId)
		if err != nil {
			log.Errorln(err)
			return err
		}

		tx, err := er.L2BridgeContract[scrollChainId].WithdrawETHtoL1(opts, er.poolAddresses[ethereumChainId], withAmount)
		if err != nil {
			log.Errorln("get bridge transaction by abi fail", "error", err)
			return err
		}

		finalTx, err := er.RawL2BridgeContract[scrollChainId].RawTransact(opts, tx.Data())
		if err != nil {
			log.Errorln("raw send bridge transaction fail", "error", err)
			return err
		}

		ctxwt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = scrollClient.SendTransaction(ctxwt, finalTx)
		if err != nil {
			log.Errorln("send bridge transaction fail", "error", err)
		}

		receipt, err := scrollClient.TxReceiptDetailByHash(finalTx.Hash())
		if err != nil {
			log.Errorln("get transaction by hash fail", "error", err)
			return err
		}

		log.Infoln("withdraw scroll pool eth to l1 pool transaction success", "txHash", receipt.TxHash)
	}

	l1LastestBlockNumber, _ := ethereumClient.GetLatestBlock()
	l2LastestBlockNumber, _ := scrollClient.GetLatestBlock()
	var l1PoolBalance *big.Int
	var l2PoolBalance *big.Int

	if l1TokenAddress == common2.ETH {
		l1PoolBalance, _ = ethereumClient.GetBalanceByBlockNumber(er.poolAddresses[ethereumChainId].String(), l1LastestBlockNumber)
		l2PoolBalance, _ = scrollClient.GetBalanceByBlockNumber(er.poolAddresses[scrollChainId].String(), l2LastestBlockNumber)

	} else {
		l1PoolBalance, _ = ethereumClient.GetERC20Balance(common.HexToAddress(l1TokenAddress), er.poolAddresses[ethereumChainId], l1LastestBlockNumber)
		l2PoolBalance, _ = scrollClient.GetERC20Balance(common.HexToAddress(l2TokenAddress), er.poolAddresses[scrollChainId], l1LastestBlockNumber)
	}

	*ethereumPoolBalance = l1PoolBalance
	*scrollPoolBalance = l2PoolBalance

	return nil
}

func (er *Exporter) metricEthereumPoolBalance() {
	ticker := time.NewTicker(60 * time.Second)
	for {
		<-ticker.C
		latestBlockNumber, err := er.ethClients[er.l1ChainID].GetLatestBlock()
		if err != nil {
			log.Errorln("get latest block number fail", "err", err)
			return
		}

		ethBalance, err := er.ethClients[er.l1ChainID].GetBalanceByBlockNumber(er.poolAddresses[er.l1ChainID].String(), latestBlockNumber)
		if err != nil {
			log.Errorln("get eth balance by block number fail", "err", err)
			return
		}
		wethBalance, err := er.ethClients[er.l1ChainID].GetERC20Balance(er.WEthAddress[er.l1ChainID], er.poolAddresses[er.l1ChainID], latestBlockNumber)
		if err != nil {
			log.Errorln("get weth balance by block number fail", "err", err)
			return
		}
		usdtBalance, err := er.ethClients[er.l1ChainID].GetERC20Balance(er.USDTAddress[er.l1ChainID], er.poolAddresses[er.l1ChainID], latestBlockNumber)
		if err != nil {
			log.Errorln("get usdt balance by block number fail", "err", err)
			return
		}

		eBF, _ := ethBalance.Float64()
		wBF, _ := wethBalance.Float64()
		uBF, _ := usdtBalance.Float64()

		ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[er.l1ChainID].String(), er.EthAddress[er.l1ChainID].String()).Set(eBF)
		ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[er.l1ChainID].String(), er.EthAddress[er.l1ChainID].String()).Set(wBF)
		ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[er.l1ChainID].String(), er.EthAddress[er.l1ChainID].String()).Set(uBF)
	}
}

func (er *Exporter) getSignOpts(chainId uint64) (*bind.TransactOpts, error) {
	var opts *bind.TransactOpts
	var err error

	if er.EnableHsm {
		seqBytes, err := hex.DecodeString(er.HsmCreden)
		if err != nil {
			log.Errorln("selaginella", "decode hsm creden fail", err.Error())
			return nil, err
		}
		apikey := option.WithCredentialsJSON(seqBytes)
		kClient, err := kms.NewKeyManagementClient(context.Background(), apikey)
		if err != nil {
			log.Errorln("selaginella", "create signer error", err.Error())
			return nil, err
		}
		mk := &sign.ManagedKey{
			KeyName:      er.HsmAPIName,
			EthereumAddr: common.HexToAddress(er.HsmAddress),
			Gclient:      kClient,
		}
		opts, err = mk.NewEthereumTransactorWithChainID(context.Background(), new(big.Int).SetUint64(chainId))
		if err != nil {
			log.Errorln("selaginella", "create signer error", err.Error())
			return nil, err
		}
	} else {
		if er.privateKey == nil {
			log.Errorln("selaginella", "create signer error", "no private key provided")
			return nil, err
		}

		opts, err = bind.NewKeyedTransactorWithChainID(er.privateKey, new(big.Int).SetUint64(chainId))
		if err != nil {
			return nil, err
		}
	}
	return opts, err
}
