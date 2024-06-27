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

type TransferL2Share struct {
	GUID              uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	StrategyAddress   common.Address `gorm:"column:strategy_address;serializer:bytes" db:"strategy_address" json:"strategy_address" form:"strategy_address"`
	SharesAmount      *big.Int       `gorm:"serializer:u256;column:shares_amount" db:"shares_amount" json:"shares_amount" form:"shares_amount"`
	ChainID           *big.Int       `gorm:"serializer:u256;column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`
	FromAddress       common.Address `gorm:"column:from_address;serializer:bytes" db:"from_address" json:"from_address" form:"from_address"`
	ToAddress         common.Address `gorm:"column:to_address;serializer:bytes" db:"to_address" json:"to_address" form:"to_address"`
	StakeMessageNonce *big.Int       `gorm:"serializer:u256;column:stake_message_nonce" db:"stake_message_nonce" json:"stake_message_nonce" form:"stake_message_nonce"`
	TxHash            common.Hash    `gorm:"column:tx_hash;serializer:bytes" db:"tx_hash" json:"tx_hash" form:"tx_hash"`
	Status            int8           `gorm:"column:status" db:"status" json:"status" form:"status"`
	Timestamp         int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (TransferL2Share) TableName() string {
	return "transfer_l2_share"
}

type TransferL2ShareDB interface {
	TransferL2ShareView
	StoreTransferL2Share([]TransferL2Share) error
	BuildTransferL2Share(in *pb.TransferL2ShareRequest) []TransferL2Share
	ChangeTransferL2ShareSentStatusByTxHash(txHash string) error
	UpdateTransferL2ShareTransactionHash(TransferL2Share) error
	UpdateTransferL2ShareFailStatus(TransferL2Share) error
}

type TransferL2ShareView interface {
	TransferL2ShareByNonce(nonce uint64) (*TransferL2Share, error)
	TransferL2ShareByTxHash(txHash string) (*TransferL2Share, error)
	OldestPendingSentTransaction() (*TransferL2Share, error)
	OldestPendingNoSentTransaction() (*TransferL2Share, error)
}

type transferL2ShareDB struct {
	gorm *gorm.DB
}

func NewTransferL2ShareDB(db *gorm.DB) TransferL2ShareDB {
	return &transferL2ShareDB{gorm: db}
}

func (t *transferL2ShareDB) BuildTransferL2Share(in *pb.TransferL2ShareRequest) []TransferL2Share {
	var transferL2Shares []TransferL2Share

	for chainID, shares := range in.ShareRequest {
		for strategy, share := range shares.ShareMap {
			sharesB, _ := new(big.Int).SetString(share, 10)
			transferL2Share := TransferL2Share{
				GUID:              uuid.New(),
				StrategyAddress:   common.HexToAddress(strategy),
				SharesAmount:      sharesB,
				ChainID:           new(big.Int).SetUint64(chainID),
				FromAddress:       common.HexToAddress(in.From),
				ToAddress:         common.HexToAddress(in.To),
				StakeMessageNonce: new(big.Int).SetUint64(in.StakeMessageNonce),
				TxHash:            common.Hash{},
				Status:            PendingStatus,
				Timestamp:         time.Now().Unix(),
			}
			transferL2Shares = append(transferL2Shares, transferL2Share)
		}
	}

	return transferL2Shares
}

func (t *transferL2ShareDB) UpdateTransferL2ShareTransactionHash(update TransferL2Share) error {
	result := t.gorm.Table("transfer_l2_share").Where("guid = ?", update.GUID.String()).Updates(map[string]interface{}{"tx_hash": update.TxHash.String()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *transferL2ShareDB) ChangeTransferL2ShareSentStatusByTxHash(txHash string) error {
	result := t.gorm.Table("transfer_l2_share").Where("tx_hash = ?", txHash).Updates(map[string]interface{}{"status": sentStatus})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *transferL2ShareDB) TransferL2ShareByTxHash(txHash string) (*TransferL2Share, error) {
	var transferL2Share TransferL2Share
	result := t.gorm.Table("transfer_l2_share").Where("tx_hash = ?", txHash).Take(&transferL2Share)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &transferL2Share, nil
}

func (t *transferL2ShareDB) TransferL2ShareByNonce(nonce uint64) (*TransferL2Share, error) {
	var transferL2Share TransferL2Share
	result := t.gorm.Table("transfer_l2_share").Where("stake_message_nonce = ?", nonce).First(&transferL2Share)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &transferL2Share, nil
}

func (t *transferL2ShareDB) ChangeTransferL2ShareStatusByTxHash(txHash string) error {

	result := t.gorm.Table("transfer_l2_share").Where("tx_hash = ?", txHash).Updates(map[string]interface{}{"status": sentStatus})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *transferL2ShareDB) UpdateTransferL2ShareFailStatus(transferL2Share TransferL2Share) error {
	result := t.gorm.Table("transfer_l2_share").Where("guid = ?", transferL2Share.GUID.String()).Updates(map[string]interface{}{"status": FailStatus})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *transferL2ShareDB) StoreTransferL2Share(updates []TransferL2Share) error {
	result := t.gorm.CreateInBatches(&updates, len(updates))
	return result.Error
}

func (t *transferL2ShareDB) OldestPendingNoSentTransaction() (*TransferL2Share, error) {
	var transferL2Share TransferL2Share
	result := t.gorm.Table("transfer_l2_share").Where("status = ? and tx_hash = ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&transferL2Share)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &transferL2Share, nil
}

func (t *transferL2ShareDB) OldestPendingSentTransaction() (*TransferL2Share, error) {
	var transferL2Share TransferL2Share
	result := t.gorm.Table("transfer_l2_share").Where("status = ? and tx_hash != ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&transferL2Share)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &transferL2Share, nil
}
