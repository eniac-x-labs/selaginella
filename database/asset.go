package database

import (
	"github.com/google/uuid"
)

type Asset struct {
	GUID      uuid.UUID `gorm:"primaryKey"          json:"guid"`
	AssetName string    `gorm:"type:varchar(30)"    json:"asset_name"`
	Contract  string    `gorm:"type:varchar(512)"   json:"contract"`
	Decimal   int       `gorm:"type:smallint"       json:"decimal"`
}

func (at Asset) TableName() string {
	return "asset"
}
