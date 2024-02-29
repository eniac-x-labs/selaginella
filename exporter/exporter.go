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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/evm-layer2/selaginella/bindings"
	common2 "github.com/evm-layer2/selaginella/common"
	"github.com/evm-layer2/selaginella/common/retry"
	"github.com/evm-layer2/selaginella/common/tasks"
	"github.com/evm-layer2/selaginella/config"
	"github.com/evm-layer2/selaginella/database"
	node "github.com/evm-layer2/selaginella/eth_client"
	"github.com/evm-layer2/selaginella/sign"
)

type Exporter struct {
	exporterConfig *config.Exporter
	db             *database.DB
	ethClients     map[uint64]node.EthClient
	*HsmConfig
	poolAddresses       map[uint64]common.Address
	EthAddress          map[uint64]common.Address
	WEthAddress         map[uint64]common.Address
	USDTAddress         map[uint64]common.Address
	USDCAddress         map[uint64]common.Address
	DAIAddress          map[uint64]common.Address
	RawL1BridgeContract *bind.BoundContract
	RawL2BridgeContract map[uint64]*bind.BoundContract
	L1BridgeContract    *bindings.L1PoolManager
	L2BridgeContract    map[uint64]*bindings.L2PoolManager
	privateKey          *ecdsa.PrivateKey
	l1ChainID           uint64
	stopped             atomic.Bool
	tasks               tasks.Group
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
	USDCAddress := make(map[uint64]common.Address)
	DAIAddress := make(map[uint64]common.Address)

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
			USDCAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].USDCAddress)
			DAIAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].DAIAddress)

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
			USDCAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].USDCAddress)
			DAIAddress[chainRpcCfg[i].ChainId] = common.HexToAddress(chainRpcCfg[i].DAIAddress)
		}

	}

	return &Exporter{
		exporterConfig:      &exporterConfig,
		db:                  db,
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
		USDCAddress:         USDCAddress,
		DAIAddress:          DAIAddress,
		l1ChainID:           l1ChainID,
		TransferMultiple:    multipleCfg,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in selaginella exporter: %w", err))
		}},
	}, nil
}

func (er *Exporter) Start(ctx context.Context) error {

	log.Infoln("exporter config", er.exporterConfig.ExporterAddress)
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

	//eFBTicker := time.NewTicker(2 * time.Hour)
	//er.tasks.Go(func() error {
	//	for range eFBTicker.C {
	//		err := er.metricEthFundBalance()
	//		if err != nil {
	//			log.Error(err.Error())
	//		}
	//	}
	//	return nil
	//})

	//wFBTicker := time.NewTicker(2 * time.Hour)
	//er.tasks.Go(func() error {
	//	for range wFBTicker.C {
	//		err := er.metricWEthFundBalance()
	//		if err != nil {
	//      	log.Error(err.Error())
	//		}
	//	}
	//	return nil
	//})
	//
	//uTFBTicker := time.NewTicker(2 * time.Hour)
	//er.tasks.Go(func() error {
	//	for range uTFBTicker.C {
	//		err := er.metricUSDTFundBalance()
	//		if err != nil {
	//			log.Error(err.Error())
	//		}
	//	}
	//	return nil
	//})
	//
	//uCFBTicker := time.NewTicker(2 * time.Hour)
	//er.tasks.Go(func() error {
	//	for range uCFBTicker.C {
	//		err := er.metricUSDCFundBalance()
	//		if err != nil {
	//			log.Error(err.Error())
	//		}
	//	}
	//	return nil
	//})
	//
	//dFBTicker := time.NewTicker(2 * time.Hour)
	//er.tasks.Go(func() error {
	//	for range dFBTicker.C {
	//		err := er.metricDAIFundBalance()
	//		if err != nil {
	//			log.Error(err.Error())
	//		}
	//	}
	//	return nil
	//})

	ePBTicker := time.NewTicker(1 * time.Minute)
	er.tasks.Go(func() error {
		for range ePBTicker.C {
			err := er.metricEthereumPoolBalance()
			if err != nil {
				log.Error(err.Error())
			}
		}
		return nil
	})

	log.Infoln("Listening on", er.exporterConfig.ExporterAddress)
	if err := http.ListenAndServe(er.exporterConfig.ExporterAddress, nil); err != nil {
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

func (er *Exporter) metricEthFundBalance() error {

	var cOpts *bind.CallOpts
	var balance *big.Int
	var ethereumPoolBalance *big.Int
	var ethereumChainId uint64
	var opPoolBalance *big.Int
	var opChainId uint64
	var polygonZkEvmPoolBalance *big.Int
	var polygonZkEvmChainId uint64
	var scrollPoolBalance *big.Int
	var scrollChainId uint64
	var chainCount uint64

	for chainId, client := range er.ethClients {
		latestBlockNumber, err := client.GetLatestBlock()
		if err != nil {
			log.Errorln("exporter", "client get latest block number fail", err)
			return err
		}

		cOpts = &bind.CallOpts{
			BlockNumber: latestBlockNumber,
			From:        crypto.PubkeyToAddress(er.privateKey.PublicKey),
		}

		if chainId == er.l1ChainID {
			balance, err = er.L1BridgeContract.FundingPoolBalance(cOpts, er.EthAddress[chainId])
		} else {
			balance, err = er.L2BridgeContract[chainId].FundingPoolBalance(cOpts, er.EthAddress[chainId])
		}
		if err != nil {
			log.Error("get bridge funding pool balance fail", "err", err)
			return err
		}

		switch chainId {
		case common2.ChainEthID, common2.ChainEthSepoliaID:
			ethereumChainId = chainId
			ethereumPoolBalance = balance
			chainCount++
		case common2.ChainOpID, common2.ChainOpSepoliaID:
			opChainId = chainId
			opPoolBalance = balance
			chainCount++
		case common2.ChainPolygonZkEvmID, common2.ChainPolygonZkEvmSepoliaID:
			polygonZkEvmChainId = chainId
			polygonZkEvmPoolBalance = balance
			chainCount++
		case common2.ChainScrollID, common2.ChainScrollSepoliaID:
			scrollChainId = chainId
			scrollPoolBalance = balance
			chainCount++
		default:
			log.Errorln("unknown chain")
		}
	}
	ChainTotalBalance := new(big.Int).Add(ethereumPoolBalance, opPoolBalance)
	ChainTotalBalance.Add(ChainTotalBalance, polygonZkEvmPoolBalance)
	ChainTotalBalance.Add(ChainTotalBalance, scrollPoolBalance)
	ChainAverageBalance := new(big.Int).Div(ChainTotalBalance, new(big.Int).SetUint64(chainCount))

	if err := er.metricOpBalance(&opPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, opChainId, ethereumChainId, er.EthAddress[ethereumChainId], er.EthAddress[opChainId]); err != nil {
		log.Errorln("metric op eth balance fail", "error", err)
		return err
	}
	if err := er.metricPolygonZkEvmBalance(&polygonZkEvmPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, polygonZkEvmChainId, ethereumChainId, er.EthAddress[ethereumChainId], er.EthAddress[polygonZkEvmChainId]); err != nil {
		log.Errorln("metric polygonZkEvm eth balance fail", "error", err)
		return err
	}
	if err := er.metricScrollBalance(&scrollPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, scrollChainId, ethereumChainId, er.EthAddress[ethereumChainId], er.EthAddress[scrollChainId]); err != nil {
		log.Errorln("metric scroll eth balance fail", "error", err)
		return err
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

	return nil
}

func (er *Exporter) metricWEthFundBalance() error {

	var cOpts *bind.CallOpts
	var balance *big.Int
	var ethereumPoolBalance *big.Int
	var ethereumChainId uint64
	var opPoolBalance *big.Int
	var opChainId uint64
	var polygonZkEvmPoolBalance *big.Int
	var polygonZkEvmChainId uint64
	var scrollPoolBalance *big.Int
	var scrollChainId uint64
	var chainCount uint64

	for chainId, client := range er.ethClients {
		latestBlockNumber, err := client.GetLatestBlock()
		if err != nil {
			log.Errorln("exporter", "client get latest block number fail", err)
			return err
		}

		cOpts = &bind.CallOpts{
			BlockNumber: latestBlockNumber,
			From:        crypto.PubkeyToAddress(er.privateKey.PublicKey),
		}

		if chainId == er.l1ChainID {
			balance, err = er.L1BridgeContract.FundingPoolBalance(cOpts, er.WEthAddress[chainId])
		} else {
			balance, err = er.L2BridgeContract[chainId].FundingPoolBalance(cOpts, er.WEthAddress[chainId])
		}
		if err != nil {
			log.Error("get bridge funding pool balance fail", "err", err)
			return err
		}

		switch chainId {
		case common2.ChainEthID, common2.ChainEthSepoliaID:
			ethereumChainId = chainId
			ethereumPoolBalance = balance
			chainCount++
		case common2.ChainOpID, common2.ChainOpSepoliaID:
			opChainId = chainId
			opPoolBalance = balance
			chainCount++
		case common2.ChainPolygonZkEvmID, common2.ChainPolygonZkEvmSepoliaID:
			polygonZkEvmChainId = chainId
			polygonZkEvmPoolBalance = balance
			chainCount++
		case common2.ChainScrollID, common2.ChainScrollSepoliaID:
			scrollChainId = chainId
			scrollPoolBalance = balance
			chainCount++
		default:
			log.Errorln("unknown chain")
		}
	}
	ChainTotalBalance := new(big.Int).Add(ethereumPoolBalance, opPoolBalance)
	ChainTotalBalance.Add(ChainTotalBalance, polygonZkEvmPoolBalance)
	ChainTotalBalance.Add(ChainTotalBalance, scrollPoolBalance)
	ChainAverageBalance := new(big.Int).Div(ChainTotalBalance, new(big.Int).SetUint64(chainCount))

	if err := er.metricOpBalance(&opPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, opChainId, ethereumChainId, er.WEthAddress[ethereumChainId], er.WEthAddress[opChainId]); err != nil {
		log.Errorln("metric op weth balance fail", "error", err)
		return err
	}
	if err := er.metricPolygonZkEvmBalance(&polygonZkEvmPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, polygonZkEvmChainId, ethereumChainId, er.WEthAddress[ethereumChainId], er.WEthAddress[polygonZkEvmChainId]); err != nil {
		log.Errorln("metric polygonZkEvm weth balance fail", "error", err)
		return err
	}
	if err := er.metricScrollBalance(&scrollPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, scrollChainId, ethereumChainId, er.WEthAddress[ethereumChainId], er.WEthAddress[scrollChainId]); err != nil {
		log.Errorln("metric scroll weth balance fail", "error", err)
		return err
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

	return nil
}

func (er *Exporter) metricUSDTFundBalance() error {

	var cOpts *bind.CallOpts
	var balance *big.Int
	var ethereumPoolBalance *big.Int
	var ethereumChainId uint64
	var opPoolBalance *big.Int
	var opChainId uint64
	var polygonZkEvmPoolBalance *big.Int
	var polygonZkEvmChainId uint64
	var scrollPoolBalance *big.Int
	var scrollChainId uint64
	var chainCount uint64

	for chainId, client := range er.ethClients {
		latestBlockNumber, err := client.GetLatestBlock()
		if err != nil {
			log.Errorln("exporter", "client get latest block number fail", err)
			return err
		}

		cOpts = &bind.CallOpts{
			BlockNumber: latestBlockNumber,
			From:        crypto.PubkeyToAddress(er.privateKey.PublicKey),
		}

		if chainId == er.l1ChainID {
			balance, err = er.L1BridgeContract.FundingPoolBalance(cOpts, er.USDTAddress[chainId])
		} else {
			balance, err = er.L2BridgeContract[chainId].FundingPoolBalance(cOpts, er.USDTAddress[chainId])
		}
		if err != nil {
			log.Error("get bridge funding pool balance fail", "err", err)
			return err
		}

		switch chainId {
		case common2.ChainEthID, common2.ChainEthSepoliaID:
			ethereumChainId = chainId
			ethereumPoolBalance = balance
			chainCount++
		case common2.ChainOpID, common2.ChainOpSepoliaID:
			opChainId = chainId
			opPoolBalance = balance
			chainCount++
		case common2.ChainPolygonZkEvmID, common2.ChainPolygonZkEvmSepoliaID:
			polygonZkEvmChainId = chainId
			polygonZkEvmPoolBalance = balance
			chainCount++
		case common2.ChainScrollID, common2.ChainScrollSepoliaID:
			scrollChainId = chainId
			scrollPoolBalance = balance
			chainCount++
		default:
			log.Errorln("unknown chain")
		}
	}
	ChainTotalBalance := new(big.Int).Add(ethereumPoolBalance, opPoolBalance)
	ChainTotalBalance.Add(ChainTotalBalance, polygonZkEvmPoolBalance)
	ChainTotalBalance.Add(ChainTotalBalance, scrollPoolBalance)
	ChainAverageBalance := new(big.Int).Div(ChainTotalBalance, new(big.Int).SetUint64(chainCount))

	if err := er.metricOpBalance(&opPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, opChainId, ethereumChainId, er.USDTAddress[ethereumChainId], er.USDTAddress[opChainId]); err != nil {
		log.Errorln("metric op usdt balance fail", "error", err)
		return err
	}
	if err := er.metricPolygonZkEvmBalance(&polygonZkEvmPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, polygonZkEvmChainId, ethereumChainId, er.USDTAddress[ethereumChainId], er.USDTAddress[polygonZkEvmChainId]); err != nil {
		log.Errorln("metric polygonZkEvm usdt balance fail", "error", err)
		return err
	}
	if err := er.metricScrollBalance(&scrollPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, scrollChainId, ethereumChainId, er.USDTAddress[ethereumChainId], er.USDTAddress[scrollChainId]); err != nil {
		log.Errorln("metric scroll usdt balance fail", "error", err)
		return err
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

	return nil
}

func (er *Exporter) metricUSDCFundBalance() error {

	var cOpts *bind.CallOpts
	var balance *big.Int
	var ethereumPoolBalance *big.Int
	var ethereumChainId uint64
	var opPoolBalance *big.Int
	var opChainId uint64
	var polygonZkEvmPoolBalance *big.Int
	var polygonZkEvmChainId uint64
	var scrollPoolBalance *big.Int
	var scrollChainId uint64
	var chainCount uint64

	for chainId, client := range er.ethClients {
		latestBlockNumber, err := client.GetLatestBlock()
		if err != nil {
			log.Errorln("exporter", "client get latest block number fail", err)
			return err
		}

		cOpts = &bind.CallOpts{
			BlockNumber: latestBlockNumber,
			From:        crypto.PubkeyToAddress(er.privateKey.PublicKey),
		}

		if chainId == er.l1ChainID {
			balance, err = er.L1BridgeContract.FundingPoolBalance(cOpts, er.USDCAddress[chainId])
		} else {
			balance, err = er.L2BridgeContract[chainId].FundingPoolBalance(cOpts, er.USDCAddress[chainId])
		}
		if err != nil {
			log.Error("get bridge funding pool balance fail", "err", err)
			return err
		}

		switch chainId {
		case common2.ChainEthID, common2.ChainEthSepoliaID:
			ethereumChainId = chainId
			ethereumPoolBalance = balance
			chainCount++
		case common2.ChainOpID, common2.ChainOpSepoliaID:
			opChainId = chainId
			opPoolBalance = balance
			chainCount++
		case common2.ChainPolygonZkEvmID, common2.ChainPolygonZkEvmSepoliaID:
			polygonZkEvmChainId = chainId
			polygonZkEvmPoolBalance = balance
			chainCount++
		case common2.ChainScrollID, common2.ChainScrollSepoliaID:
			scrollChainId = chainId
			scrollPoolBalance = balance
			chainCount++
		default:
			log.Errorln("unknown chain")
		}
	}
	ChainTotalBalance := new(big.Int).Add(ethereumPoolBalance, opPoolBalance)
	ChainTotalBalance.Add(ChainTotalBalance, polygonZkEvmPoolBalance)
	ChainTotalBalance.Add(ChainTotalBalance, scrollPoolBalance)
	ChainAverageBalance := new(big.Int).Div(ChainTotalBalance, new(big.Int).SetUint64(chainCount))

	if err := er.metricOpBalance(&opPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, opChainId, ethereumChainId, er.USDCAddress[ethereumChainId], er.USDCAddress[opChainId]); err != nil {
		log.Errorln("metric op usdc balance fail", "error", err)
		return err
	}
	if err := er.metricPolygonZkEvmBalance(&polygonZkEvmPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, polygonZkEvmChainId, ethereumChainId, er.USDCAddress[ethereumChainId], er.USDCAddress[polygonZkEvmChainId]); err != nil {
		log.Errorln("metric polygonZkEvm usdc balance fail", "error", err)
		return err
	}
	if err := er.metricScrollBalance(&scrollPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, scrollChainId, ethereumChainId, er.USDCAddress[ethereumChainId], er.USDCAddress[scrollChainId]); err != nil {
		log.Errorln("metric scroll usdc balance fail", "error", err)
		return err
	}

	ethPBF, _ := ethereumPoolBalance.Float64()
	opPBF, _ := opPoolBalance.Float64()
	pyPBF, _ := polygonZkEvmPoolBalance.Float64()
	scPBF, _ := scrollPoolBalance.Float64()
	averageBalanceF, _ := ChainAverageBalance.Float64()

	ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[ethereumChainId].String(), er.USDCAddress[ethereumChainId].String()).Set(ethPBF)
	opPoolBalanceMetric.WithLabelValues(er.poolAddresses[opChainId].String(), er.USDCAddress[opChainId].String()).Set(opPBF)
	polygonZkEvmPoolBalanceMetric.WithLabelValues(er.poolAddresses[polygonZkEvmChainId].String(), er.USDCAddress[polygonZkEvmChainId].String()).Set(pyPBF)
	scrollPoolBalanceMetric.WithLabelValues(er.poolAddresses[scrollChainId].String(), er.USDCAddress[scrollChainId].String()).Set(scPBF)
	chainAverageBalanceMetric.WithLabelValues("USDC").Set(averageBalanceF)

	return nil
}

func (er *Exporter) metricDAIFundBalance() error {

	var cOpts *bind.CallOpts
	var balance *big.Int
	var ethereumPoolBalance *big.Int
	var ethereumChainId uint64
	var opPoolBalance *big.Int
	var opChainId uint64
	var polygonZkEvmPoolBalance *big.Int
	var polygonZkEvmChainId uint64
	var scrollPoolBalance *big.Int
	var scrollChainId uint64
	var chainCount uint64

	for chainId, client := range er.ethClients {
		latestBlockNumber, err := client.GetLatestBlock()
		if err != nil {
			log.Errorln("exporter", "client get latest block number fail", err)
			return err
		}

		cOpts = &bind.CallOpts{
			BlockNumber: latestBlockNumber,
			From:        crypto.PubkeyToAddress(er.privateKey.PublicKey),
		}

		if chainId == er.l1ChainID {
			balance, err = er.L1BridgeContract.FundingPoolBalance(cOpts, er.DAIAddress[chainId])
		} else {
			balance, err = er.L2BridgeContract[chainId].FundingPoolBalance(cOpts, er.DAIAddress[chainId])
		}
		if err != nil {
			log.Error("get bridge funding pool balance fail", "err", err)
			return err
		}

		switch chainId {
		case common2.ChainEthID, common2.ChainEthSepoliaID:
			ethereumChainId = chainId
			ethereumPoolBalance = balance
			chainCount++
		case common2.ChainOpID, common2.ChainOpSepoliaID:
			opChainId = chainId
			opPoolBalance = balance
			chainCount++
		case common2.ChainPolygonZkEvmID, common2.ChainPolygonZkEvmSepoliaID:
			polygonZkEvmChainId = chainId
			polygonZkEvmPoolBalance = balance
			chainCount++
		case common2.ChainScrollID, common2.ChainScrollSepoliaID:
			scrollChainId = chainId
			scrollPoolBalance = balance
			chainCount++
		default:
			log.Errorln("unknown chain")
		}
	}
	ChainTotalBalance := new(big.Int).Add(ethereumPoolBalance, opPoolBalance)
	ChainTotalBalance.Add(ChainTotalBalance, polygonZkEvmPoolBalance)
	ChainTotalBalance.Add(ChainTotalBalance, scrollPoolBalance)
	ChainAverageBalance := new(big.Int).Div(ChainTotalBalance, new(big.Int).SetUint64(chainCount))

	if err := er.metricOpBalance(&opPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, opChainId, ethereumChainId, er.DAIAddress[ethereumChainId], er.DAIAddress[opChainId]); err != nil {

		log.Errorln("metric op dai balance fail", "error", err)
		return err
	}
	if err := er.metricPolygonZkEvmBalance(&polygonZkEvmPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, polygonZkEvmChainId, ethereumChainId, er.DAIAddress[ethereumChainId], er.DAIAddress[polygonZkEvmChainId]); err != nil {
		log.Errorln("metric polygonZkEvm dai balance fail", "error", err)
		return err
	}
	if err := er.metricScrollBalance(&scrollPoolBalance, &ethereumPoolBalance, ChainAverageBalance, chainCount, scrollChainId, ethereumChainId, er.DAIAddress[ethereumChainId], er.DAIAddress[scrollChainId]); err != nil {
		log.Errorln("metric scroll dai balance fail", "error", err)
		return err
	}

	ethPBF, _ := ethereumPoolBalance.Float64()
	opPBF, _ := opPoolBalance.Float64()
	pyPBF, _ := polygonZkEvmPoolBalance.Float64()
	scPBF, _ := scrollPoolBalance.Float64()
	averageBalanceF, _ := ChainAverageBalance.Float64()

	ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[ethereumChainId].String(), er.DAIAddress[ethereumChainId].String()).Set(ethPBF)
	opPoolBalanceMetric.WithLabelValues(er.poolAddresses[opChainId].String(), er.DAIAddress[opChainId].String()).Set(opPBF)
	polygonZkEvmPoolBalanceMetric.WithLabelValues(er.poolAddresses[polygonZkEvmChainId].String(), er.DAIAddress[polygonZkEvmChainId].String()).Set(pyPBF)
	scrollPoolBalanceMetric.WithLabelValues(er.poolAddresses[scrollChainId].String(), er.DAIAddress[scrollChainId].String()).Set(scPBF)
	chainAverageBalanceMetric.WithLabelValues("DAI").Set(averageBalanceF)

	return nil
}

func (er *Exporter) metricOpBalance(opPoolBalance **big.Int, ethereumPoolBalance **big.Int, ChainAverageBalance *big.Int, TotalChainNum uint64, opChainId uint64, ethereumChainId uint64, l1TokenAddress common.Address, l2TokenAddress common.Address) error {

	transferAmount := new(big.Int).SetUint64(0)
	withdrawAmount := new(big.Int).SetUint64(0)
	ethereumClient := er.ethClients[ethereumChainId]
	opClient := er.ethClients[opChainId]

	if (*opPoolBalance).Cmp(new(big.Int).Div(ChainAverageBalance, new(big.Int).SetUint64(er.L1Multiple))) < 1 {
		transferAmount = new(big.Int).Div(*ethereumPoolBalance, new(big.Int).Div(new(big.Int).SetUint64(TotalChainNum), new(big.Int).SetUint64(2)))

		receipt, err := er.transferAssertToBridge(ethereumChainId, opChainId, transferAmount, l2TokenAddress, ethereumClient)
		if err != nil {
			log.Errorln(err)
			return err
		}

		log.Infoln("send l1 pool eth to op pool transaction success", "txHash", receipt.TxHash)

	}

	if (*opPoolBalance).Cmp(new(big.Int).Mul(ChainAverageBalance, new(big.Int).SetUint64(er.L2Multiple))) > -1 {
		withdrawAmount = new(big.Int).Div(*opPoolBalance, new(big.Int).SetUint64(3))

		receipt, err := er.WithdrawToL1(opChainId, ethereumChainId, withdrawAmount, l2TokenAddress, opClient)
		if err != nil {
			log.Errorln(err)
			return err
		}

		log.Infoln("withdraw op pool eth to l1 pool transaction success", "txHash", receipt.TxHash)
	}

	var l1PoolBalance *big.Int
	var l2PoolBalance *big.Int

	l1PoolBalance = new(big.Int).Sub(*ethereumPoolBalance, transferAmount)
	l1PoolBalance = new(big.Int).Add(l1PoolBalance, withdrawAmount)
	l2PoolBalance = new(big.Int).Add(*opPoolBalance, transferAmount)
	l2PoolBalance = new(big.Int).Sub(l2PoolBalance, withdrawAmount)

	*ethereumPoolBalance = l1PoolBalance
	*opPoolBalance = l2PoolBalance

	return nil
}

func (er *Exporter) metricPolygonZkEvmBalance(pyPoolBalance **big.Int, ethereumPoolBalance **big.Int, ChainAverageBalance *big.Int, TotalChainNum uint64, pyChainId uint64, ethereumChainId uint64, l1TokenAddress common.Address, l2TokenAddress common.Address) error {

	transferAmount := new(big.Int).SetUint64(0)
	withdrawAmount := new(big.Int).SetUint64(0)
	ethereumClient := er.ethClients[ethereumChainId]
	pyClient := er.ethClients[pyChainId]

	if (*pyPoolBalance).Cmp(new(big.Int).Div(ChainAverageBalance, new(big.Int).SetUint64(er.L1Multiple))) < 1 {
		transferAmount = new(big.Int).Div(*ethereumPoolBalance, new(big.Int).Div(new(big.Int).SetUint64(TotalChainNum), new(big.Int).SetUint64(2)))

		receipt, err := er.transferAssertToBridge(ethereumChainId, pyChainId, transferAmount, l2TokenAddress, ethereumClient)
		if err != nil {
			log.Errorln(err)
			return err
		}

		log.Infoln("send l1 pool eth to polygon zkEvm pool transaction success", "txHash", receipt.TxHash)

	}

	if (*pyPoolBalance).Cmp(new(big.Int).Mul(ChainAverageBalance, new(big.Int).SetUint64(er.L2Multiple))) > -1 {
		withdrawAmount = new(big.Int).Div(*pyPoolBalance, new(big.Int).SetUint64(3))

		receipt, err := er.WithdrawToL1(pyChainId, ethereumChainId, withdrawAmount, l2TokenAddress, pyClient)
		if err != nil {
			log.Errorln(err)
			return err
		}
		log.Infoln("withdraw polygon zkEvm pool eth to l1 pool transaction success", "txHash", receipt.TxHash)
	}

	var l1PoolBalance *big.Int
	var l2PoolBalance *big.Int

	l1PoolBalance = new(big.Int).Sub(*ethereumPoolBalance, transferAmount)
	l1PoolBalance = new(big.Int).Add(l1PoolBalance, withdrawAmount)
	l2PoolBalance = new(big.Int).Add(*pyPoolBalance, transferAmount)
	l2PoolBalance = new(big.Int).Sub(l2PoolBalance, withdrawAmount)

	*ethereumPoolBalance = l1PoolBalance
	*pyPoolBalance = l2PoolBalance

	return nil
}

func (er *Exporter) metricScrollBalance(scrollPoolBalance **big.Int, ethereumPoolBalance **big.Int, ChainAverageBalance *big.Int, TotalChainNum uint64, scrollChainId uint64, ethereumChainId uint64, l1TokenAddress common.Address, l2TokenAddress common.Address) error {

	transferAmount := new(big.Int).SetUint64(0)
	withdrawAmount := new(big.Int).SetUint64(0)
	ethereumClient := er.ethClients[ethereumChainId]
	scrollClient := er.ethClients[scrollChainId]

	if (*scrollPoolBalance).Cmp(new(big.Int).Div(ChainAverageBalance, new(big.Int).SetUint64(er.L1Multiple))) < 1 {
		transferAmount = new(big.Int).Div(*ethereumPoolBalance, new(big.Int).Div(new(big.Int).SetUint64(TotalChainNum), new(big.Int).SetUint64(2)))

		receipt, err := er.transferAssertToBridge(ethereumChainId, scrollChainId, transferAmount, l1TokenAddress, ethereumClient)
		if err != nil {
			log.Errorln(err)
			return err
		}

		log.Infoln("send l1 pool eth to scroll pool transaction success", "txHash", receipt.TxHash)

	}

	if (*scrollPoolBalance).Cmp(new(big.Int).Mul(ChainAverageBalance, new(big.Int).SetUint64(er.L2Multiple))) > -1 {
		withdrawAmount = new(big.Int).Div(*scrollPoolBalance, new(big.Int).SetUint64(3))

		receipt, err := er.WithdrawToL1(scrollChainId, ethereumChainId, withdrawAmount, l2TokenAddress, scrollClient)
		if err != nil {
			log.Errorln(err)
			return err
		}

		log.Infoln("withdraw scroll pool eth to l1 pool transaction success", "txHash", receipt.TxHash)
	}

	var l1PoolBalance *big.Int
	var l2PoolBalance *big.Int

	l1PoolBalance = new(big.Int).Sub(*ethereumPoolBalance, transferAmount)
	l1PoolBalance = new(big.Int).Add(l1PoolBalance, withdrawAmount)
	l2PoolBalance = new(big.Int).Add(*scrollPoolBalance, transferAmount)
	l2PoolBalance = new(big.Int).Sub(l2PoolBalance, withdrawAmount)

	*ethereumPoolBalance = l1PoolBalance
	*scrollPoolBalance = l2PoolBalance

	return nil
}

func (er *Exporter) metricEthereumPoolBalance() error {

	latestBlockNumber, err := er.ethClients[er.l1ChainID].GetLatestBlock()
	if err != nil {
		log.Errorln("get latest block number fail", "err", err)
		return err
	}

	cOpts := &bind.CallOpts{
		BlockNumber: latestBlockNumber,
		From:        crypto.PubkeyToAddress(er.privateKey.PublicKey),
	}

	ethBalance, err := er.L1BridgeContract.FundingPoolBalance(cOpts, er.EthAddress[er.l1ChainID])
	if err != nil {
		log.Errorln("get eth balance by block number fail", "err", err)
		return err
	}
	wethBalance, err := er.L1BridgeContract.FundingPoolBalance(cOpts, er.WEthAddress[er.l1ChainID])
	if err != nil {
		log.Errorln("get weth balance by block number fail", "err", err)
		return err
	}
	usdtBalance, err := er.L1BridgeContract.FundingPoolBalance(cOpts, er.USDTAddress[er.l1ChainID])
	if err != nil {
		log.Errorln("get usdt balance by block number fail", "err", err)
		return err
	}
	usdcBalance, err := er.L1BridgeContract.FundingPoolBalance(cOpts, er.USDCAddress[er.l1ChainID])
	if err != nil {
		log.Errorln("get usdc balance by block number fail", "err", err)
		return err
	}
	daiBalance, err := er.L1BridgeContract.FundingPoolBalance(cOpts, er.DAIAddress[er.l1ChainID])
	if err != nil {
		log.Errorln("get dai balance by block number fail", "err", err)
		return err
	}

	eBF, _ := ethBalance.Float64()
	wBF, _ := wethBalance.Float64()
	uTBF, _ := usdtBalance.Float64()
	uCBF, _ := usdcBalance.Float64()
	dBF, _ := daiBalance.Float64()

	ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[er.l1ChainID].String(), er.EthAddress[er.l1ChainID].String()).Set(eBF)
	ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[er.l1ChainID].String(), er.WEthAddress[er.l1ChainID].String()).Set(wBF)
	ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[er.l1ChainID].String(), er.USDTAddress[er.l1ChainID].String()).Set(uTBF)
	ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[er.l1ChainID].String(), er.USDCAddress[er.l1ChainID].String()).Set(uCBF)
	ethereumPoolBalanceMetric.WithLabelValues(er.poolAddresses[er.l1ChainID].String(), er.DAIAddress[er.l1ChainID].String()).Set(dBF)

	log.Infoln("get l1 pool balance success")
	return nil
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

	opts.Context = context.Background()
	opts.NoSend = true

	return opts, err
}

func getTransactionReceipt(client node.EthClient, tx types.Transaction) (*types.Receipt, error) {
	var receipt *types.Receipt
	var err error
	var ctx = context.Background()

	retryStrategy := &retry.ExponentialStrategy{Min: 30_000_000_000, Max: 60_000_000_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](ctx, 10, retryStrategy, func() (interface{}, error) {
		receipt, err = client.TxReceiptDetailByHash(tx.Hash())
		if err != nil && !errors.Is(err, ethereum.NotFound) {
			log.Errorln("get transaction by hash fail", "error", err)
			return nil, err
		}
		return nil, nil
	}); err != nil {
		return nil, err
	}

	return receipt, nil
}

func (er *Exporter) transferAssertToBridge(ethereumChainId uint64, l2ChainId uint64, transferAmount *big.Int, tokenAddress common.Address, ethereumClient node.EthClient) (*types.Receipt, error) {
	var receipt *types.Receipt

	opts, err := er.getSignOpts(ethereumChainId)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	tx, err := er.L1BridgeContract.TransferAssertToBridge(opts, new(big.Int).SetUint64(l2ChainId), tokenAddress, er.poolAddresses[l2ChainId], transferAmount)
	if err != nil {
		log.Errorln("get bridge transaction by abi fail", "error", err)
		return nil, err
	}

	finalTx, err := er.RawL1BridgeContract.RawTransact(opts, tx.Data())
	if err != nil {
		log.Errorln("raw send bridge transaction fail", "error", err)
		return nil, err
	}

	ctxwt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = ethereumClient.SendTransaction(ctxwt, finalTx)
	if err != nil {
		log.Errorln("send bridge transaction fail", "error", err)
		return nil, err
	}

	receipt, err = getTransactionReceipt(ethereumClient, *finalTx)
	if err != nil {
		log.Errorln("get bridge transaction receipt fail", "error", err)
		return nil, err
	}

	return receipt, nil
}

func (er *Exporter) WithdrawToL1(l2ChainId uint64, ethereumChainId uint64, withdrawAmount *big.Int, tokenAddress common.Address, l2Client node.EthClient) (*types.Receipt, error) {
	var receipt *types.Receipt
	var tx *types.Transaction

	opts, err := er.getSignOpts(l2ChainId)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	switch tokenAddress.String() {
	case er.EthAddress[l2ChainId].String():
		tx, err = er.L2BridgeContract[l2ChainId].WithdrawETHtoL1(opts, er.poolAddresses[ethereumChainId], withdrawAmount)
		if err != nil {
			log.Errorln("get bridge transaction by abi fail", "error", err)
			return nil, err
		}
	case er.WEthAddress[l2ChainId].String():
		tx, err = er.L2BridgeContract[l2ChainId].WithdrawWETHToL1(opts, er.poolAddresses[ethereumChainId], withdrawAmount)
		if err != nil {
			log.Errorln("get bridge transaction by abi fail", "error", err)
			return nil, err
		}
	default:
		tx, err = er.L2BridgeContract[l2ChainId].WithdrawERC20ToL1(opts, tokenAddress, er.poolAddresses[ethereumChainId], withdrawAmount)
		if err != nil {
			log.Errorln("get bridge transaction by abi fail", "error", err)
			return nil, err
		}
	}

	finalTx, err := er.RawL2BridgeContract[l2ChainId].RawTransact(opts, tx.Data())
	if err != nil {
		log.Errorln("raw send bridge transaction fail", "error", err)
		return nil, err
	}

	ctxwt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = l2Client.SendTransaction(ctxwt, finalTx)
	if err != nil {
		log.Errorln("send bridge transaction fail", "error", err)
	}

	receipt, err = getTransactionReceipt(l2Client, *finalTx)
	if err != nil {
		log.Errorln("get bridge transaction receipt fail", "error", err)
		return nil, err
	}

	return receipt, nil
}
