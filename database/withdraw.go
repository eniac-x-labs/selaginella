package database

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Withdraws struct {
	GUID         uuid.UUID       `gorm:"primaryKey"                json:"guid"`
	AssetId      uint64          `gorm:"index"                     json:"asset_id"`
	WalletUuid   string          `gorm:"index;type:varchar(256)"   json:"wallet_uuid"`
	UserUuid     string          `gorm:"type:varchar(256)"         json:"user_uuid"`
	WithdrawId   uint64          `gorm:"index"                     json:"withdraw_id"`
	FromAddr     string          `gorm:"type:varchar(100)"         json:"from_addr"`
	ToAddr       string          `gorm:"type:varchar(100)"         json:"to_addr"`
	Amount       decimal.Decimal `sql:"type:decimal(32,18)"        json:"amount"`
	TxHash       string          `gorm:"type:varchar(100)"         json:"tx_hash"`
	RawTx        string          `gorm:"type:text"                 json:"raw_tx"`
	Status       int8            `gorm:"type:smallint"             json:"status"`         // 0: Not issued; 1: Issued; 2: Withdrawal successful; 3: Withdrawal failed; 4: Reported to the business layer successfully 5: Reported to the business layer failed
	ReceiveTxFee decimal.Decimal `sql:"type:decimal(32,18)"        json:"receive_tx_fee"` // Handling fee given by the business party
	ChainTxFee   decimal.Decimal `sql:"type:decimal(32,18)"        json:"chain_tx_fee"`   // Actual handling fees consumed on the chain
}

func (wd Withdraws) TableName() string {
	return "withdraws"
}
