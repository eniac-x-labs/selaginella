package database

import (
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"

	"github.com/evm-layer2/selaginella/protobuf/pb"
)

const (
	PendingStatus = 0
	sentStatus    = 1
	SuccessStatus = 2
)

type CrossChainTransfer struct {
	GUID                uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	SourceChainId       *big.Int       `gorm:"serializer:u256;column:source_chain_id" db:"source_chain_id" json:"source_chain_id" form:"source_chain_id"`
	DestChainId         *big.Int       `gorm:"serializer:u256;column:dest_chain_id" db:"dest_chain_id" json:"dest_chain_id" form:"dest_chain_id"`
	Fee                 *big.Int       `gorm:"serializer:u256;column:fee" db:"fee" json:"fee" form:"fee"`
	Nonce               *big.Int       `gorm:"serializer:u256;column:nonce" db:"nonce" json:"nonce" form:"nonce"`
	SourceHash          common.Hash    `gorm:"column:source_hash;serializer:bytes" db:"source_hash" json:"source_hash" form:"source_hash"`
	TxHash              common.Hash    `gorm:"column:tx_hash;serializer:bytes" db:"tx_hash" json:"tx_hash" form:"tx_hash"`
	SourceSenderAddress common.Address `gorm:"column:source_sender_address;serializer:bytes" db:"source_sender_address" json:"source_sender_address" form:"source_sender_address"`
	DestReceiveAddress  common.Address `gorm:"column:dest_receive_address;serializer:bytes" db:"dest_receive_address" json:"dest_receive_address" form:"dest_receive_address"`
	TokenAddress        common.Address `gorm:"column:token_address;serializer:bytes" db:"token_address" json:"token_address" form:"token_address"`
	Amount              *big.Int       `gorm:"serializer:u256;column:amount" db:"amount" json:"amount" form:"amount"`
	Status              int8           `gorm:"column:status" db:"status" json:"status" form:"status"`
	Timestamp           int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (CrossChainTransfer) TableName() string {
	return "cross_chain_transfer"
}

type CrossChainTransferDB interface {
	CrossChainTransferView
	StoreBatchCrossChainTransfer([]CrossChainTransfer) error
	BuildCrossChainTransfer(in *pb.CrossChainTransferRequest, sourceHash common.Hash) CrossChainTransfer
	CrossChainTransferBySourceHash(sourceHash string) (*CrossChainTransfer, error)
	ChangeCrossChainTransferSentStatueByTxHash(txHash string) error
	ChangeCrossChainTransferSuccessStatueByTxHash(txHash string) error
	UpdateCrossChainTransferTransactionHash(CrossChainTransfer) error
}

type CrossChainTransferView interface {
	CrossChainTransferByTxHash(txHash string) (*CrossChainTransfer, error)
	OldestPendingSentTransaction() (*CrossChainTransfer, error)
	OldestPendingNoSentTransaction() (*CrossChainTransfer, error)
	GetPeriodTotalFee(startTimestamp uint64, endTimeStamp uint64, tokenAddress common.Address) (*big.Int, error)
}

type crossChainTransferDB struct {
	gorm *gorm.DB
}

func NewCrossChainTransferDB(db *gorm.DB) CrossChainTransferDB {
	return &crossChainTransferDB{gorm: db}
}

func (c *crossChainTransferDB) CrossChainTransferByTxHash(txHash string) (*CrossChainTransfer, error) {
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

func (c *crossChainTransferDB) CrossChainTransferBySourceHash(sourceHash string) (*CrossChainTransfer, error) {
	var crossChainTransfer CrossChainTransfer
	result := c.gorm.Table("cross_chain_transfer").Where("source_hash = ?", sourceHash).Take(&crossChainTransfer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &crossChainTransfer, nil
}

func (c *crossChainTransferDB) ChangeCrossChainTransferSentStatueByTxHash(txHash string) error {
	var crossChainTransfer CrossChainTransfer
	result := c.gorm.Table("cross_chain_transfer").Where("tx_hash = ?", txHash).Take(&crossChainTransfer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	crossChainTransfer.Status = sentStatus
	err := c.gorm.Save(&crossChainTransfer).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *crossChainTransferDB) ChangeCrossChainTransferSuccessStatueByTxHash(txHash string) error {
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

func (c *crossChainTransferDB) StoreBatchCrossChainTransfer(transfers []CrossChainTransfer) error {
	result := c.gorm.CreateInBatches(&transfers, len(transfers))
	return result.Error
}

func (c *crossChainTransferDB) BuildCrossChainTransfer(in *pb.CrossChainTransferRequest, sourceHash common.Hash) CrossChainTransfer {

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
		SourceHash:          sourceHash,
		TxHash:              common.Hash{},
		SourceSenderAddress: common.Address{},
		DestReceiveAddress:  common.HexToAddress(in.ReceiveAddress),
		TokenAddress:        common.HexToAddress(in.TokenAddress),
		Amount:              amount,
		Status:              PendingStatus,
		Timestamp:           time.Now().Unix(),
	}
}

func (c *crossChainTransferDB) OldestPendingNoSentTransaction() (*CrossChainTransfer, error) {
	var crossChainTransfer CrossChainTransfer
	result := c.gorm.Table("cross_chain_transfer").Where("status = ? and tx_hash = ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&crossChainTransfer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &crossChainTransfer, nil
}

func (c *crossChainTransferDB) OldestPendingSentTransaction() (*CrossChainTransfer, error) {
	var crossChainTransfer CrossChainTransfer
	result := c.gorm.Table("cross_chain_transfer").Where("status = ? and tx_hash != ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&crossChainTransfer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &crossChainTransfer, nil
}

func (c *crossChainTransferDB) UpdateCrossChainTransferTransactionHash(transfer CrossChainTransfer) error {
	result := c.gorm.Table("cross_chain_transfer").Where("guid = ?", transfer.GUID.String()).Updates(map[string]interface{}{"tx_hash": transfer.TxHash.String()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *crossChainTransferDB) GetPeriodTotalFee(startTimestamp uint64, endTimeStamp uint64, tokenAddress common.Address) (*big.Int, error) {
	var totalFee string
	err := c.gorm.Table("cross_chain_transfer").Where("timestamp >= ? and timestamp < ? and token_address = ?", startTimestamp, endTimeStamp, strings.ToLower(tokenAddress.String())).Select("SUM(fee) as total_fee").Row().Scan(&totalFee)
	if err != nil {
		return nil, err
	}
	if totalFee == "" {
		return new(big.Int).SetUint64(0), nil
	}

	totalFeeB, _ := new(big.Int).SetString(totalFee, 10)
	return totalFeeB, nil
}
