package database

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	GUID    uuid.UUID       `gorm:"primaryKey"           json:"guid"`
	BlockId uint64          `gorm:"type:integer"         json:"blocks_id"`
	TxHash  string          `gorm:"type:varchar(100)"    json:"tx_hash"`
	Amount  decimal.Decimal `sql:"type:decimal(32,18)"   json:"quantity"`
}

func (tt Transaction) TableName() string {
	return "transaction"
}
