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

const (
	PendingStatus = 0
	SuccessStatus = 1
)

type CrossChainTransfer struct {
	GUID                uuid.UUID      `gorm:"primaryKey" json:"guid"`
	SourceChainId       *big.Int       `gorm:"serializer:u256;column:source_chain_id" db:"source_chain_id" json:"source_chain_id" form:"source_chain_id"`
	DestChainId         *big.Int       `gorm:"serializer:u256;column:dest_chain_id" db:"dest_chain_id" json:"dest_chain_id" form:"dest_chain_id"`
	Fee                 *big.Int       `gorm:"serializer:u256;column:fee" db:"fee" json:"fee" form:"fee"`
	Nonce               *big.Int       `gorm:"serializer:u256;column:nonce" db:"nonce" json:"nonce" form:"nonce"`
	TxHash              common.Hash    `gorm:"column:tx_hash;serializer:bytes" db:"tx_hash" json:"tx_hash" form:"tx_hash"`
	SourceSenderAddress common.Address `gorm:"column:source_sender_dddress;serializer:bytes" db:"source_sender_dddress" json:"source_sender_dddress" form:"source_sender_dddress"`
	DestReceiveAddress  common.Address `gorm:"column:dest_receive_address;serializer:bytes" db:"dest_receive_address" json:"dest_receive_address" form:"dest_receive_address"`
	TokenAddress        common.Address `gorm:"column:token_address;serializer:bytes" db:"token_address" json:"token_address" form:"token_address"`
	Amount              *big.Int       `gorm:"serializer:u256;column:amount" db:"amount" json:"amount" form:"amount"`
	Status              int8           `gorm:"serializer:u256;column:status" db:"status" json:"status" form:"status"`
	Timestamp           int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (CrossChainTransfer) TableName() string {
	return "cross_chain_transfer"
}

type CrossChainTransferDB interface {
	CrossChainTransferView
	StoreBatchCrossChainTransfer([]CrossChainTransfer) error
	BuildCrossChainTransfer(in *pb.CrossChainTransferRequest, txHash common.Hash) CrossChainTransfer
	ChangeCrossChainTransferStatueByTxHash(txHash string) error
}

type CrossChainTransferView interface {
	CrossChainTransferByTxHash(txHash string) (*CrossChainTransfer, error)
}

type crossChainTransferDB struct {
	gorm *gorm.DB
}

func NewCrossChainTransferDB(db *gorm.DB) CrossChainTransferDB {
	return &crossChainTransferDB{gorm: db}
}

func (c crossChainTransferDB) CrossChainTransferByTxHash(txHash string) (*CrossChainTransfer, error) {
	var crossChainTransfer CrossChainTransfer
	result := c.gorm.Table("cross_chain_transfer").Where("tx_hash = ?", txHash).Take(&crossChainTransfer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &crossChainTransfer, nil
}

func (c crossChainTransferDB) ChangeCrossChainTransferStatueByTxHash(txHash string) error {
	var crossChainTransfer CrossChainTransfer
	result := c.gorm.Table("cross_chain_transfer").Where("tx_hash = ?", txHash).Take(&crossChainTransfer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	crossChainTransfer.Status = SuccessStatus
	err := c.gorm.Save(&crossChainTransfer).Error
	if err != nil {
		return err
	}

	return nil
}

func (c crossChainTransferDB) StoreBatchCrossChainTransfer(transfers []CrossChainTransfer) error {
	result := c.gorm.CreateInBatches(&transfers, len(transfers))
	return result.Error
}

func (c crossChainTransferDB) BuildCrossChainTransfer(in *pb.CrossChainTransferRequest, txHash common.Hash) CrossChainTransfer {

	sci, _ := new(big.Int).SetString(in.SourceChainId, 10)
	dci, _ := new(big.Int).SetString(in.DestChainId, 10)
	amount, _ := new(big.Int).SetString(in.Amount, 10)
	fee, _ := new(big.Int).SetString(in.Fee, 10)
	nonce, _ := new(big.Int).SetString(in.Nonce, 10)

	return CrossChainTransfer{
		GUID:                uuid.New(),
		SourceChainId:       sci,
		DestChainId:         dci,
		Fee:                 fee,
		Nonce:               nonce,
		TxHash:              txHash,
		SourceSenderAddress: common.Address{},
		DestReceiveAddress:  common.HexToAddress(in.ReceiveAddress),
		TokenAddress:        common.HexToAddress(in.TokenAddress),
		Amount:              amount,
		Status:              PendingStatus,
		Timestamp:           time.Now().Unix(),
	}
}
