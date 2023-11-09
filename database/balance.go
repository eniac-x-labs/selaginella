package database

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Balance struct {
	GUID      uuid.UUID       `gorm:"primaryKey"                    json:"guid"`
	AddressId uint64          `gorm:"index"                         json:"address_id"`
	AssetID   uint64          `gorm:"index"                         json:"asset_id"`
	Type      uint            `gorm:"type:smallint"                 json:"type"` //0:system 1:wallet
	Amount    decimal.Decimal `sql:"type:decimal(32,18);default:0"  json:"quantity"`
}

func (bl Balance) TableName() string {
	return "balance"
}
