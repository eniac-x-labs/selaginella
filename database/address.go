package database

import "github.com/google/uuid"

type Address struct {
	GUID       uuid.UUID `gorm:"primaryKey"                json:"guid"`
	WalletUuid uuid.UUID `gorm:"index;type:varchar(256)"   json:"wallet_uuid"`
	UserUuid   uuid.UUID `gorm:"type:varchar(256)"         json:"user_uuid"`
	Type       uint      `gorm:"type:smallint"             json:"type"` // 0:normal address 1:fee collection; 2:cold wallet address
	Address    string    `gorm:"type:varchar(126)"         json:"address"`
	PrivateKey string    `gorm:"type:varchar(512)"         json:"private_key"`
}

func (addr Address) TableName() string {
	return "address"
}
