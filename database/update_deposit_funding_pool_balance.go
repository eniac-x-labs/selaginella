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

type UpdateDepositFundingPoolBalance struct {
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

func (UpdateDepositFundingPoolBalance) TableName() string {
	return "update_deposit_funding_pool_balance"
}

type UpdateDepositFundingPoolBalanceDB interface {
	UpdateDepositFundingPoolBalanceView
	StoreBatchUpdateFundingPoolBalance([]UpdateDepositFundingPoolBalance) error
	BuildUpdateFundingPoolBalance(in *pb.UpdateDepositFundingPoolBalanceRequest, sourceHash common.Hash) UpdateDepositFundingPoolBalance
	UpdateFundingPoolBalanceBySourceHash(sourceHash string) (*UpdateDepositFundingPoolBalance, error)
	ChangeUpdateFundingPoolBalanceSentStatueByTxHash(txHash string) error
	UpdateFundingPoolBalanceTransactionHash(UpdateDepositFundingPoolBalance) error
}

type UpdateDepositFundingPoolBalanceView interface {
	UpdateFundingPoolBalanceByTxHash(txHash string) (*UpdateDepositFundingPoolBalance, error)
	OldestPendingSentTransaction() (*UpdateDepositFundingPoolBalance, error)
	OldestPendingNoSentTransaction() (*UpdateDepositFundingPoolBalance, error)
}

type updateDepositFundingPoolBalanceDB struct {
	gorm *gorm.DB
}

func (u *updateDepositFundingPoolBalanceDB) BuildUpdateFundingPoolBalance(in *pb.UpdateDepositFundingPoolBalanceRequest, sourceHash common.Hash) UpdateDepositFundingPoolBalance {
	sci, _ := new(big.Int).SetString(in.SourceChainId, 10)
	dci, _ := new(big.Int).SetString(in.DestChainId, 10)
	amount, _ := new(big.Int).SetString(in.Amount, 10)

	return UpdateDepositFundingPoolBalance{
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

func (u *updateDepositFundingPoolBalanceDB) UpdateFundingPoolBalanceTransactionHash(update UpdateDepositFundingPoolBalance) error {
	result := u.gorm.Table("update_deposit_funding_pool_balance").Where("guid = ?", update.GUID.String()).Updates(map[string]interface{}{"tx_hash": update.TxHash.String()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewDepositUpdateFundingPoolBalanceDB(db *gorm.DB) UpdateDepositFundingPoolBalanceDB {
	return &updateDepositFundingPoolBalanceDB{gorm: db}
}

func (u *updateDepositFundingPoolBalanceDB) UpdateFundingPoolBalanceByTxHash(txHash string) (*UpdateDepositFundingPoolBalance, error) {
	var updateFundingPoolBalance UpdateDepositFundingPoolBalance
	result := u.gorm.Table("update_deposit_funding_pool_balance").Where("tx_hash = ?", txHash).Take(&updateFundingPoolBalance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &updateFundingPoolBalance, nil
}

func (u *updateDepositFundingPoolBalanceDB) UpdateFundingPoolBalanceBySourceHash(sourceHash string) (*UpdateDepositFundingPoolBalance, error) {
	var updateFundingPoolBalance UpdateDepositFundingPoolBalance
	result := u.gorm.Table("update_deposit_funding_pool_balance").Where("source_hash = ?", sourceHash).Take(&updateFundingPoolBalance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &updateFundingPoolBalance, nil
}

func (u *updateDepositFundingPoolBalanceDB) ChangeUpdateFundingPoolBalanceSentStatueByTxHash(txHash string) error {
	var updateFundingPoolBalance UpdateDepositFundingPoolBalance
	result := u.gorm.Table("update_deposit_funding_pool_balance").Where("tx_hash = ?", txHash).Take(&updateFundingPoolBalance)
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

func (u *updateDepositFundingPoolBalanceDB) StoreBatchUpdateFundingPoolBalance(updates []UpdateDepositFundingPoolBalance) error {
	result := u.gorm.CreateInBatches(&updates, len(updates))
	return result.Error
}

func (u *updateDepositFundingPoolBalanceDB) OldestPendingNoSentTransaction() (*UpdateDepositFundingPoolBalance, error) {
	var updateFundingPoolBalance UpdateDepositFundingPoolBalance
	result := u.gorm.Table("update_deposit_funding_pool_balance").Where("status = ? and tx_hash = ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&updateFundingPoolBalance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &updateFundingPoolBalance, nil
}

func (u *updateDepositFundingPoolBalanceDB) OldestPendingSentTransaction() (*UpdateDepositFundingPoolBalance, error) {
	var updateFundingPoolBalance UpdateDepositFundingPoolBalance
	result := u.gorm.Table("update_deposit_funding_pool_balance").Where("status = ? and tx_hash != ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&updateFundingPoolBalance)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &updateFundingPoolBalance, nil
}
