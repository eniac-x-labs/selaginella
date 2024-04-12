package database

import (
	"errors"
	"math/big"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"

	"github.com/evm-layer2/selaginella/protobuf/pb"
)

type TransferToL2Bridge struct {
	GUID            uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	Batch           *big.Int       `gorm:"serializer:u256;column:batch" db:"batch" json:"batch" form:"batch"`
	ChainId         *big.Int       `gorm:"serializer:u256;column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`
	StrategyAddress common.Address `gorm:"column:strategy_address;serializer:bytes" db:"strategy_address" json:"strategy_address" form:"strategy_address"`
	TxHash          common.Hash    `gorm:"column:tx_hash;serializer:bytes" db:"tx_hash" json:"tx_hash" form:"tx_hash"`
	Status          int8           `gorm:"column:status" db:"status" json:"status" form:"status"`
	Timestamp       int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (TransferToL2Bridge) TableName() string {
	return "transfer_to_l2_bridge"
}

type TransferToL2BridgeDB interface {
	TransferToL2BridgeView
	StoreTransferToL2Bridge([]TransferToL2Bridge) error
	BuildTransferToL2Bridge(in *pb.TransferToL2DappLinkBridgeRequest) TransferToL2Bridge
	ChangeTransferToL2BridgeSentStatusByTxHash(txHash string) error
	UpdateTransferToL2BridgeTransactionHash(TransferToL2Bridge) error
}

type TransferToL2BridgeView interface {
	TransferToL2BridgeByBatch(batch uint64) (*TransferToL2Bridge, error)
	TransferToL2BridgeByTxHash(txHash string) (*TransferToL2Bridge, error)
	OldestPendingSentTransaction() (*TransferToL2Bridge, error)
	OldestPendingNoSentTransaction() (*TransferToL2Bridge, error)
}

type transferToL2BridgeDB struct {
	gorm *gorm.DB
}

func NewTransferToL2BridgeDB(db *gorm.DB) TransferToL2BridgeDB {
	return &transferToL2BridgeDB{gorm: db}
}

func (t *transferToL2BridgeDB) BuildTransferToL2Bridge(in *pb.TransferToL2DappLinkBridgeRequest) TransferToL2Bridge {

	ci, _ := new(big.Int).SetString(in.ChainId, 10)

	return TransferToL2Bridge{
		GUID:            uuid.New(),
		Batch:           new(big.Int).SetUint64(in.Batch),
		ChainId:         ci,
		StrategyAddress: common.HexToAddress(in.StrategyAddress),
		TxHash:          common.Hash{},
		Status:          PendingStatus,
		Timestamp:       time.Now().Unix(),
	}
}

func (t *transferToL2BridgeDB) UpdateTransferToL2BridgeTransactionHash(update TransferToL2Bridge) error {

	result := t.gorm.Table("transfer_to_l2_bridge").Where("guid = ?", update.GUID).Updates(map[string]interface{}{"tx_hash": update.TxHash.String()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *transferToL2BridgeDB) TransferToL2BridgeByTxHash(txHash string) (*TransferToL2Bridge, error) {
	var transferToL2Bridge TransferToL2Bridge
	result := t.gorm.Table("transfer_to_l2_bridge").Where("tx_hash = ?", txHash).Take(&transferToL2Bridge)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &transferToL2Bridge, nil
}

func (t *transferToL2BridgeDB) TransferToL2BridgeByBatch(batch uint64) (*TransferToL2Bridge, error) {
	var transferToL2Bridge TransferToL2Bridge
	result := t.gorm.Table("transfer_to_l2_bridge").Where("batch = ?", batch).Take(&transferToL2Bridge)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &transferToL2Bridge, nil
}

func (t *transferToL2BridgeDB) ChangeTransferToL2BridgeSentStatusByTxHash(txHash string) error {

	result := t.gorm.Table("transfer_to_l2_bridge").Where("tx_hash = ?", txHash).Updates(map[string]interface{}{"status": sentStatus})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *transferToL2BridgeDB) StoreTransferToL2Bridge(updates []TransferToL2Bridge) error {
	result := t.gorm.CreateInBatches(&updates, len(updates))
	return result.Error
}

func (t *transferToL2BridgeDB) OldestPendingNoSentTransaction() (*TransferToL2Bridge, error) {
	var transferToL2Bridge TransferToL2Bridge
	result := t.gorm.Table("transfer_to_l2_bridge").Where("status = ? and tx_hash = ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&transferToL2Bridge)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &transferToL2Bridge, nil
}

func (t *transferToL2BridgeDB) OldestPendingSentTransaction() (*TransferToL2Bridge, error) {
	var transferToL2Bridge TransferToL2Bridge
	result := t.gorm.Table("transfer_to_l2_bridge").Where("status = ? and tx_hash != ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&transferToL2Bridge)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &transferToL2Bridge, nil
}
