package db

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Deposits struct {
	GUID       uuid.UUID       `gorm:"primaryKey"               json:"guid"`
	WalletUuid string          `gorm:"index;type:varchar(256)"  json:"wallet_uuid"`
	UserUuid   string          `gorm:"type:varchar(256)"        json:"user_uuid"`
	AssetId    uint64          `gorm:"index"                    json:"asset_id"`
	FromAddr   string          `gorm:"type:varchar(100)"        json:"from_addr"`
	ToAddr     string          `gorm:"type:varchar(100)"        json:"to_addr"`
	Amount     decimal.Decimal `sql:"type:decimal(32,18)"       json:"amount"`
	TxHash     string          `gorm:"type:varchar(100)"        json:"tx_hash"`
	TxFee      decimal.Decimal `sql:"type:decimal(32,18)"       json:"tx_fee"`
	Status     int8            `gorm:"type:smallint"            json:"status"` //0:deposit success but do not notify business，1:notify business success, 2:notify business fail
}

func (dt Deposits) TableName() string {
	return "deposits"
}
