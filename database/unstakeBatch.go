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

type UnstakeBatch struct {
	GUID            uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	StrategyAddress common.Address `gorm:"column:strategy_address;serializer:bytes" db:"strategy_address" json:"strategy_address" form:"strategy_address"`
	BridgeAddress   common.Address `gorm:"column:bridge_address;serializer:bytes" db:"bridge_address" json:"bridge_address" form:"bridge_address"`
	SourceChainId   *big.Int       `gorm:"serializer:u256;column:source_chain_id" db:"source_chain_id" json:"source_chain_id" form:"source_chain_id"`
	DestChainId     *big.Int       `gorm:"serializer:u256;column:dest_chain_id" db:"dest_chain_id" json:"dest_chain_id" form:"dest_chain_id"`
	TxHash          common.Hash    `gorm:"column:tx_hash;serializer:bytes" db:"tx_hash" json:"tx_hash" form:"tx_hash"`
	GasLimit        *big.Int       `gorm:"serializer:u256;column:gas_limit" db:"gas_limit" json:"gas_limit" form:"gas_limit"`
	SourceHash      common.Hash    `gorm:"column:source_hash;serializer:bytes" db:"source_hash" json:"source_hash" form:"source_hash"`
	Status          int8           `gorm:"column:status" db:"status" json:"status" form:"status"`
	Timestamp       int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (UnstakeBatch) TableName() string {
	return "unstake_batch"
}

type UnstakeBatchDB interface {
	UnstakeBatchView
	StoreUnstakeBatch([]UnstakeBatch) error
	BuildUnstakeBatch(in *pb.UnstakeBatchRequest, sourceHash common.Hash) UnstakeBatch
	UnstakeBatchBySourceHash(sourceHash string) (*UnstakeBatch, error)
	ChangeUnstakeBatchSentStatusByTxHash(txHash string) error
	UpdateUnstakeBatchTransactionHash(UnstakeBatch) error
}

type UnstakeBatchView interface {
	UnstakeBatchByTxHash(txHash string) (*UnstakeBatch, error)
	OldestPendingSentTransaction() (*UnstakeBatch, error)
	OldestPendingNoSentTransaction() (*UnstakeBatch, error)
}

type unstakeBatchDB struct {
	gorm *gorm.DB
}

func NewUnstakeBatchDB(db *gorm.DB) UnstakeBatchDB {
	return &unstakeBatchDB{gorm: db}
}

func (u *unstakeBatchDB) BuildUnstakeBatch(in *pb.UnstakeBatchRequest, sourceHash common.Hash) UnstakeBatch {
	sci, _ := new(big.Int).SetString(in.SourceChainId, 10)
	dci, _ := new(big.Int).SetString(in.DestChainId, 10)
	gasLimit, _ := new(big.Int).SetString(in.GasLimit, 10)

	return UnstakeBatch{
		GUID:            uuid.New(),
		StrategyAddress: common.HexToAddress(in.L2StrategyAddress),
		BridgeAddress:   common.HexToAddress(in.BridgeAddress),
		SourceChainId:   sci,
		DestChainId:     dci,
		TxHash:          common.Hash{},
		GasLimit:        gasLimit,
		SourceHash:      sourceHash,
		Status:          PendingStatus,
		Timestamp:       time.Now().Unix(),
	}
}

func (u *unstakeBatchDB) UpdateUnstakeBatchTransactionHash(update UnstakeBatch) error {
	result := u.gorm.Table("unstake_batch").Where("guid = ?", update.GUID.String()).Updates(map[string]interface{}{"tx_hash": update.TxHash.String()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *unstakeBatchDB) UnstakeBatchByTxHash(txHash string) (*UnstakeBatch, error) {
	var unstakeBatch UnstakeBatch
	result := u.gorm.Table("unstake_batch").Where("tx_hash = ?", txHash).Take(&unstakeBatch)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &unstakeBatch, nil
}

func (u *unstakeBatchDB) UnstakeBatchBySourceHash(sourceHash string) (*UnstakeBatch, error) {
	var unstakeBatch UnstakeBatch
	result := u.gorm.Table("unstake_batch").Where("source_hash = ?", sourceHash).Take(&unstakeBatch)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &unstakeBatch, nil
}

func (u *unstakeBatchDB) ChangeUnstakeBatchSentStatusByTxHash(txHash string) error {
	var unstakeBatch UnstakeBatch
	result := u.gorm.Table("unstake_batch").Where("tx_hash = ?", txHash).Take(&unstakeBatch)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	unstakeBatch.Status = sentStatus
	err := u.gorm.Save(&unstakeBatch).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *unstakeBatchDB) StoreUnstakeBatch(updates []UnstakeBatch) error {
	result := u.gorm.CreateInBatches(&updates, len(updates))
	return result.Error
}

func (u *unstakeBatchDB) OldestPendingNoSentTransaction() (*UnstakeBatch, error) {
	var unstakeBatch UnstakeBatch
	result := u.gorm.Table("unstake_batch").Where("status = ? and tx_hash = ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&unstakeBatch)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &unstakeBatch, nil
}

func (u *unstakeBatchDB) OldestPendingSentTransaction() (*UnstakeBatch, error) {
	var unstakeBatch UnstakeBatch
	result := u.gorm.Table("unstake_batch").Where("status = ? and tx_hash != ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&unstakeBatch)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &unstakeBatch, nil
}
