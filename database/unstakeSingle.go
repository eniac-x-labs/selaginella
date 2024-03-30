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

type UnstakeSingle struct {
	GUID            uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	StakerAddress   common.Address `gorm:"column:staker_address;serializer:bytes" db:"staker_address" json:"staker_address" form:"staker_address"`
	StrategyAddress common.Address `gorm:"column:strategy_address;serializer:bytes" db:"strategy_address" json:"strategy_address" form:"strategy_address"`
	SharesAmount    *big.Int       `gorm:"serializer:u256;column:shares_amount" db:"shares_amount" json:"shares_amount" form:"shares_amount"`
	ChainId         *big.Int       `gorm:"serializer:u256;column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`
	TxHash          common.Hash    `gorm:"column:tx_hash;serializer:bytes" db:"tx_hash" json:"tx_hash" form:"tx_hash"`
	SourceHash      common.Hash    `gorm:"column:source_hash;serializer:bytes" db:"source_hash" json:"source_hash" form:"source_hash"`
	Status          int8           `gorm:"column:status" db:"status" json:"status" form:"status"`
	Timestamp       int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (UnstakeSingle) TableName() string {
	return "unstake_single"
}

type UnstakeSingleDB interface {
	UnstakeSingleView
	StoreUnstakeSingle([]UnstakeSingle) error
	BuildUnstakeSingle(in *pb.UnstakeSingleRequest, sourceHash common.Hash) UnstakeSingle
	UnstakeSingleBySourceHash(SourceHash string) (*UnstakeSingle, error)
	ChangeUnstakeSingleSentStatusByTxHash(txHash string) error
	UpdateUnstakeSingleTransactionHash(UnstakeSingle) error
}

type UnstakeSingleView interface {
	UnstakeSingleByTxHash(txHash string) (*UnstakeSingle, error)
	OldestPendingSentTransaction() (*UnstakeSingle, error)
	OldestPendingNoSentTransaction() (*UnstakeSingle, error)
}

type unstakeSingleDB struct {
	gorm *gorm.DB
}

func NewUnstakeSingleDB(db *gorm.DB) UnstakeSingleDB {
	return &unstakeSingleDB{gorm: db}
}

func (u *unstakeSingleDB) BuildUnstakeSingle(in *pb.UnstakeSingleRequest, sourceHash common.Hash) UnstakeSingle {
	sa, _ := new(big.Int).SetString(in.SharesAmount, 10)
	ci, _ := new(big.Int).SetString(in.ChainId, 10)

	return UnstakeSingle{
		GUID:            uuid.New(),
		StakerAddress:   common.HexToAddress(in.StakerAddress),
		StrategyAddress: common.HexToAddress(in.StrategyAddress),
		SharesAmount:    sa,
		ChainId:         ci,
		TxHash:          common.Hash{},
		SourceHash:      sourceHash,
		Status:          PendingStatus,
		Timestamp:       time.Now().Unix(),
	}
}

func (u *unstakeSingleDB) UpdateUnstakeSingleTransactionHash(update UnstakeSingle) error {
	result := u.gorm.Table("unstake_single").Where("guid = ?", update.GUID.String()).Updates(map[string]interface{}{"tx_hash": update.TxHash.String()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *unstakeSingleDB) UnstakeSingleByTxHash(txHash string) (*UnstakeSingle, error) {
	var unstakeSingle UnstakeSingle
	result := u.gorm.Table("unstake_single").Where("tx_hash = ?", txHash).Take(&unstakeSingle)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &unstakeSingle, nil
}

func (u *unstakeSingleDB) UnstakeSingleBySourceHash(sourceHash string) (*UnstakeSingle, error) {
	var unstakeSingle UnstakeSingle
	result := u.gorm.Table("unstake_single").Where("source_hash = ?", sourceHash).Take(&unstakeSingle)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &unstakeSingle, nil
}

func (u *unstakeSingleDB) ChangeUnstakeSingleSentStatusByTxHash(txHash string) error {
	var unstakeSingle UnstakeSingle
	result := u.gorm.Table("unstake_single").Where("tx_hash = ?", txHash).Take(&unstakeSingle)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	unstakeSingle.Status = sentStatus
	err := u.gorm.Save(&unstakeSingle).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *unstakeSingleDB) StoreUnstakeSingle(updates []UnstakeSingle) error {
	result := u.gorm.CreateInBatches(&updates, len(updates))
	return result.Error
}

func (u *unstakeSingleDB) OldestPendingNoSentTransaction() (*UnstakeSingle, error) {
	var unstakeSingle UnstakeSingle
	result := u.gorm.Table("unstake_single").Where("status = ? and tx_hash = ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&unstakeSingle)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &unstakeSingle, nil
}

func (u *unstakeSingleDB) OldestPendingSentTransaction() (*UnstakeSingle, error) {
	var unstakeSingle UnstakeSingle
	result := u.gorm.Table("unstake_single").Where("status = ? and tx_hash != ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&unstakeSingle)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &unstakeSingle, nil
}
