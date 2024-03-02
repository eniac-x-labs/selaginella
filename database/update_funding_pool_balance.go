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

type UpdateFundingPoolBalance struct {
	GUID           uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	SourceChainId  *big.Int       `gorm:"serializer:u256;column:source_chain_id" db:"source_chain_id" json:"source_chain_id" form:"source_chain_id"`
	DestChainId    *big.Int       `gorm:"serializer:u256;column:dest_chain_id" db:"dest_chain_id" json:"dest_chain_id" form:"dest_chain_id"`
	SourceHash     common.Hash    `gorm:"column:source_hash;serializer:bytes" db:"source_hash" json:"source_hash" form:"source_hash"`
	TxHash         common.Hash    `gorm:"column:tx_hash;serializer:bytes" db:"tx_hash" json:"tx_hash" form:"tx_hash"`
	ReceiveAddress common.Address `gorm:"column:receive_address;serializer:bytes" db:"receive_address" json:"receive_address" form:"receive_address"`
	TokenAddress   common.Address `gorm:"column:token_address;serializer:bytes" db:"token_address" json:"token_address" form:"token_address"`
	Amount         *big.Int       `gorm:"serializer:u256;column:amount" db:"amount" json:"amount" form:"amount"`
	Status         int8           `gorm:"column:status" db:"status" json:"status" form:"status"`
	Timestamp      int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (UpdateFundingPoolBalance) TableName() string {
	return "update_funding_pool_balance"
}

type UpdateFundingPoolBalanceDB interface {
	UpdateFundingPoolBalanceView
	StoreBatchUpdateFundingPoolBalance([]UpdateFundingPoolBalance) error
	BuildUpdateFundingPoolBalance(in *pb.UpdateFundingPoolBalanceRequest, sourceHash common.Hash) UpdateFundingPoolBalance
	UpdateFundingPoolBalanceBySourceHash(sourceHash string) (*UpdateFundingPoolBalance, error)
	ChangeUpdateFundingPoolBalanceSentStatueByTxHash(txHash string) error
	UpdateFundingPoolBalanceTransactionHash(UpdateFundingPoolBalance) error
}

type UpdateFundingPoolBalanceView interface {
	UpdateFundingPoolBalanceByTxHash(txHash string) (*UpdateFundingPoolBalance, error)
	OldestPendingSentTransaction() (*UpdateFundingPoolBalance, error)
	OldestPendingNoSentTransaction() (*UpdateFundingPoolBalance, error)
}

type updateFundingPoolBalanceDB struct {
	gorm *gorm.DB
}

func (u *updateFundingPoolBalanceDB) BuildUpdateFundingPoolBalance(in *pb.UpdateFundingPoolBalanceRequest, sourceHash common.Hash) UpdateFundingPoolBalance {
	sci, _ := new(big.Int).SetString(in.SourceChainId, 10)
	dci, _ := new(big.Int).SetString(in.DestChainId, 10)
	amount, _ := new(big.Int).SetString(in.Amount, 10)

	return UpdateFundingPoolBalance{
		GUID:           uuid.New(),
		SourceChainId:  sci,
		DestChainId:    dci,
		SourceHash:     sourceHash,
		TxHash:         common.Hash{},
		ReceiveAddress: common.HexToAddress(in.ReceiveAddress),
		TokenAddress:   common.HexToAddress(in.TokenAddress),
		Amount:         amount,
		Status:         PendingStatus,
		Timestamp:      time.Now().Unix(),
	}
}

func (u *updateFundingPoolBalanceDB) UpdateFundingPoolBalanceTransactionHash(update UpdateFundingPoolBalance) error {
	result := u.gorm.Table("update_funding_pool_balance").Where("guid = ?", update.GUID.String()).Updates(map[string]interface{}{"tx_hash": update.TxHash.String()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewUpdateFundingPoolBalanceDB(db *gorm.DB) UpdateFundingPoolBalanceDB {
	return &updateFundingPoolBalanceDB{gorm: db}
}

func (u *updateFundingPoolBalanceDB) UpdateFundingPoolBalanceByTxHash(txHash string) (*UpdateFundingPoolBalance, error) {
	var updateFundingPoolBalance UpdateFundingPoolBalance
	result := u.gorm.Table("update_funding_pool_balance").Where("tx_hash = ?", txHash).Take(&updateFundingPoolBalance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &updateFundingPoolBalance, nil
}

func (u *updateFundingPoolBalanceDB) UpdateFundingPoolBalanceBySourceHash(sourceHash string) (*UpdateFundingPoolBalance, error) {
	var updateFundingPoolBalance UpdateFundingPoolBalance
	result := u.gorm.Table("update_funding_pool_balance").Where("source_hash = ?", sourceHash).Take(&updateFundingPoolBalance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &updateFundingPoolBalance, nil
}

func (u *updateFundingPoolBalanceDB) ChangeUpdateFundingPoolBalanceSentStatueByTxHash(txHash string) error {
	var updateFundingPoolBalance UpdateFundingPoolBalance
	result := u.gorm.Table("update_funding_pool_balance").Where("tx_hash = ?", txHash).Take(&updateFundingPoolBalance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	updateFundingPoolBalance.Status = sentStatus
	err := u.gorm.Save(&updateFundingPoolBalance).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *updateFundingPoolBalanceDB) StoreBatchUpdateFundingPoolBalance(updates []UpdateFundingPoolBalance) error {
	result := u.gorm.CreateInBatches(&updates, len(updates))
	return result.Error
}

func (u *updateFundingPoolBalanceDB) OldestPendingNoSentTransaction() (*UpdateFundingPoolBalance, error) {
	var updateFundingPoolBalance UpdateFundingPoolBalance
	result := u.gorm.Table("update_funding_pool_balance").Where("status = ? and tx_hash = ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&updateFundingPoolBalance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &updateFundingPoolBalance, nil
}

func (u *updateFundingPoolBalanceDB) OldestPendingSentTransaction() (*UpdateFundingPoolBalance, error) {
	var updateFundingPoolBalance UpdateFundingPoolBalance
	result := u.gorm.Table("update_funding_pool_balance").Where("status = ? and tx_hash != ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&updateFundingPoolBalance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &updateFundingPoolBalance, nil
}
