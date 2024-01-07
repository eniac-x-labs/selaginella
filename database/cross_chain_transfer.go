package database

import (
	"github.com/google/uuid"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type CrossChainTransfer struct {
	GUID                uuid.UUID      `gorm:"primaryKey" json:"guid"`
	SourceChainId       *big.Int       `gorm:"serializer:u256;column:source_chain_id" db:"source_chain_id" json:"source_chain_id" form:"source_chain_id"`
	DestChainId         *big.Int       `gorm:"serializer:u256;column:dest_chain_id" db:"dest_chain_id" json:"dest_chain_id" form:"dest_chain_id"`
	Fee                 *big.Int       `gorm:"serializer:u256;column:fee" db:"fee" json:"fee" form:"fee"`
	TxHash              common.Hash    `gorm:"column:tx_hash;serializer:bytes" db:"tx_hash" json:"tx_hash" form:"tx_hash"`
	SourceSenderAddress common.Address `gorm:"column:source_sender_dddress;serializer:bytes" db:"source_sender_dddress" json:"source_sender_dddress" form:"source_sender_dddress"`
	DestReceiveAddress  common.Address `gorm:"column:dest_receive_address;serializer:bytes" db:"dest_receive_address" json:"dest_receive_address" form:"dest_receive_address"`
	Amount              *big.Int       `gorm:"serializer:u256;column:amount" db:"amount" json:"amount" form:"amount"`
	Status              int8           `gorm:"serializer:u256;column:status" db:"status" json:"status" form:"status"`
	Timestamp           int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (CrossChainTransfer) TableName() string {
	return "cross_chain_transfer"
}

type CrossChainTransferDB interface {
	CrossChainTransferView
	StoreBatchDataStores([]CrossChainTransfer) error
}

type CrossChainTransferView interface {
	DataStoreBlockById(txHash common.Hash) (*CrossChainTransfer, error)
}

type crossChainTransferDB struct {
	gorm *gorm.DB
}

func NewCrossChainTransferDB(db *gorm.DB) CrossChainTransferDB {
	return &crossChainTransferDB{gorm: db}
}

func (c crossChainTransferDB) DataStoreBlockById(txHash common.Hash) (*CrossChainTransfer, error) {
	panic("implement me")
}

func (c crossChainTransferDB) StoreBatchDataStores(transfers []CrossChainTransfer) error {
	panic("implement me")
}
