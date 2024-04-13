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

type BatchMint struct {
	GUID          uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	StakerAddress common.Address `gorm:"column:staker_address;serializer:bytes" db:"staker_address" json:"staker_address" form:"staker_address"`
	SharesAmount  *big.Int       `gorm:"serializer:u256;column:shares_amount" db:"shares_amount" json:"shares_amount" form:"shares_amount"`
	Batch         *big.Int       `gorm:"serializer:u256;column:batch" db:"batch" json:"batch" form:"batch"`
	TxHash        common.Hash    `gorm:"column:tx_hash;serializer:bytes" db:"tx_hash" json:"tx_hash" form:"tx_hash"`
	Status        int8           `gorm:"column:status" db:"status" json:"status" form:"status"`
	Timestamp     int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (BatchMint) TableName() string {
	return "batch_mint"
}

type BatchMintDB interface {
	BatchMintView
	StoreBatchMint([]BatchMint) error
	BuildBatchMint(in *pb.BatchMintRequest) []BatchMint
	ChangeBatchMintSentStatusByTxHash(txHash string) error
	UpdateBatchMintTransactionHash([]BatchMint) error
	UpdateBatchMintFailStatus(BatchMint) error
}

type BatchMintView interface {
	BatchMintByBatch(batch uint64) (*BatchMint, error)
	BatchMintsByBatch(batch uint64) ([]BatchMint, error)
	BatchMintByTxHash(txHash string) (*BatchMint, error)
	OldestPendingSentTransaction() (*BatchMint, error)
	OldestPendingNoSentTransaction() (*BatchMint, error)
}

type batchMintDB struct {
	gorm *gorm.DB
}

func NewBatchMintDB(db *gorm.DB) BatchMintDB {
	return &batchMintDB{gorm: db}
}

func (b *batchMintDB) BuildBatchMint(in *pb.BatchMintRequest) []BatchMint {
	var batches []BatchMint

	for addr, shares := range in.Mint {
		sharesB, _ := new(big.Int).SetString(shares, 10)
		batch := BatchMint{
			GUID:          uuid.New(),
			StakerAddress: common.HexToAddress(addr),
			SharesAmount:  sharesB,
			Batch:         new(big.Int).SetUint64(in.Batch),
			TxHash:        common.Hash{},
			Status:        PendingStatus,
			Timestamp:     time.Now().Unix(),
		}
		batches = append(batches, batch)
	}

	return batches
}

func (b *batchMintDB) UpdateBatchMintTransactionHash(update []BatchMint) error {
	batch := update[0].Batch
	txHash := update[0].TxHash
	result := b.gorm.Table("batch_mint").Where("batch = ?", batch).Updates(map[string]interface{}{"tx_hash": txHash.String()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *batchMintDB) BatchMintByTxHash(txHash string) (*BatchMint, error) {
	var batchMint BatchMint
	result := b.gorm.Table("batch_mint").Where("tx_hash = ?", txHash).Take(&batchMint)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &batchMint, nil
}

func (b *batchMintDB) BatchMintByBatch(batch uint64) (*BatchMint, error) {
	var batchMint BatchMint
	result := b.gorm.Table("batch_mint").Where("batch = ?", batch).First(&batchMint)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &batchMint, nil
}

func (b *batchMintDB) BatchMintsByBatch(batch uint64) ([]BatchMint, error) {
	var batchMints []BatchMint
	result := b.gorm.Table("batch_mint").Where("batch = ?", batch).Find(&batchMints)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return batchMints, nil
}

func (b *batchMintDB) ChangeBatchMintSentStatusByTxHash(txHash string) error {

	result := b.gorm.Table("batch_mint").Where("tx_hash = ?", txHash).Updates(map[string]interface{}{"status": sentStatus})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *batchMintDB) UpdateBatchMintFailStatus(batch BatchMint) error {
	result := b.gorm.Table("batch_mint").Where("batch = ?", batch.Batch).Updates(map[string]interface{}{"status": FailStatus})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *batchMintDB) StoreBatchMint(updates []BatchMint) error {
	result := b.gorm.CreateInBatches(&updates, len(updates))
	return result.Error
}

func (b *batchMintDB) OldestPendingNoSentTransaction() (*BatchMint, error) {
	var batchMint BatchMint
	result := b.gorm.Table("batch_mint").Where("status = ? and tx_hash = ?", PendingStatus, common.Hash{}.String()).Order("batch asc").Find(&batchMint)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &batchMint, nil
}

func (b *batchMintDB) OldestPendingSentTransaction() (*BatchMint, error) {
	var batchMint BatchMint
	result := b.gorm.Table("batch_mint").Where("status = ? and tx_hash != ?", PendingStatus, common.Hash{}.String()).Order("batch asc").Find(&batchMint)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &batchMint, nil
}
