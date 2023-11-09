package database

import "github.com/google/uuid"

type Blocks struct {
	GUID      uuid.UUID `gorm:"primaryKey"         json:"guid"`
	Height    uint64    `gorm:"type:integer"       json:"height;"             `
	BlockHash string    `gorm:"type:varchar(100)"  json:"block_hash"`
}

func (bl Blocks) TableName() string {
	return "blocks"
}
