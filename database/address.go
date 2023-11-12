package database

import (
	"gorm.io/gorm"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/google/uuid"
)

type Address struct {
	GUID       uuid.UUID `gorm:"primaryKey"                json:"guid"`
	WalletGuid uuid.UUID `gorm:"index;type:varchar(256)"   json:"wallet_guid"`
	UserGuid   uuid.UUID `gorm:"type:varchar(256)"         json:"user_guid"`
	Type       uint      `gorm:"type:smallint"             json:"type"` // 0:normal address 1:fee collection; 2:cold wallet address
	Address    string    `gorm:"type:varchar(126)"         json:"address"`
	PrivateKey string    `gorm:"type:varchar(512)"         json:"private_key"`
}

func (Address) TableName() string {
	return "address"
}

type AddressDB interface {
	AddressView
	StoreAddressTransactions([]Address) error
}

type AddressView interface {
	GetAddressInfo(string) (Address, error)
}

/**
 * Implementation
 */

type addressDB struct {
	gorm *gorm.DB
}

func (a addressDB) GetAddressInfo(s string) (Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a addressDB) StoreAddressTransactions(addresses []Address) error {
	//TODO implement me
	panic("implement me")
}

func NewAddressDB(db *gorm.DB) AddressDB {
	gplus.Init(db)
	return &addressDB{gorm: db}
}
