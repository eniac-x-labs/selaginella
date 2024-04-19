package database

import (
	"errors"
	"math/big"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"

	"github.com/evm-layer2/selaginella/protobuf/pb"
)

const (
	mSentStatus = 1
	qSentStatus = 2
)

type MigrateL1Shares struct {
	GUID              uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')" json:"guid"`
	UnstakerAddress   common.Address `gorm:"column:unstaker_address;serializer:bytes" db:"unstaker_address" json:"unstaker_address" form:"unstaker_address"`
	StrategyAddress   common.Address `gorm:"column:strategy_address;serializer:bytes" db:"strategy_address" json:"strategy_address" form:"strategy_address"`
	SharesAmount      *big.Int       `gorm:"serializer:u256;column:shares_amount" db:"shares_amount" json:"shares_amount" form:"shares_amount"`
	L1UnstakeMsgNonce *big.Int       `gorm:"serializer:u256;column:l1_unstake_msg_nonce" db:"l1_unstake_msg_nonce" json:"l1_unstake_msg_nonce" form:"l1_unstake_msg_nonce"`
	ChainId           *big.Int       `gorm:"serializer:u256;column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`
	MTxHash           common.Hash    `gorm:"column:m_tx_hash;serializer:bytes" db:"m_tx_hash" json:"m_tx_hash" form:"m_tx_hash"`
	QTxHash           common.Hash    `gorm:"column:q_tx_hash;serializer:bytes" db:"q_tx_hash" json:"q_tx_hash" form:"q_tx_hash"`
	SourceHash        common.Hash    `gorm:"column:source_hash;serializer:bytes" db:"source_hash" json:"source_hash" form:"source_hash"`
	Status            int8           `gorm:"column:status" db:"status" json:"status" form:"status"`
	Timestamp         int64          `gorm:"column:timestamp" db:"timestamp" json:"timestamp" form:"timestamp"`
}

func (MigrateL1Shares) TableName() string {
	return "migrate_l1_shares"
}

type MigrateL1SharesDB interface {
	MigrateL1SharesView
	StoreMigrateL1Shares([]MigrateL1Shares) error
	BuildMigrateL1Shares(in *pb.MigrateL1SharesRequest, sourceHash common.Hash) MigrateL1Shares
	MigrateL1SharesBySourceHash(SourceHash string) (*MigrateL1Shares, error)
	ChangeMigrateL1SharesSentStatusByTxHash(txHash string) error
	ChangeQueueWithdrawalsSentStatusByTxHash(txHash string) error
	UpdateMigrateL1SharesTransactionHash(MigrateL1Shares) error
	UpdateQueueWithdrawalsTransactionHash(update MigrateL1Shares) error
}

type MigrateL1SharesView interface {
	MigrateL1SharesByTxHash(txHash string) (*MigrateL1Shares, error)
	OldestMPendingSentTransaction() (*MigrateL1Shares, error)
	OldestQPendingSentTransaction() (*MigrateL1Shares, error)
	OldestPendingNoSentTransaction() (*MigrateL1Shares, error)
}

type migrateL1SharesDB struct {
	gorm *gorm.DB
}

func NewMigrateL1SharesDB(db *gorm.DB) MigrateL1SharesDB {
	return &migrateL1SharesDB{gorm: db}
}

func (m *migrateL1SharesDB) BuildMigrateL1Shares(in *pb.MigrateL1SharesRequest, sourceHash common.Hash) MigrateL1Shares {
	ci, _ := new(big.Int).SetString(in.ChainId, 10)
	share, _ := new(big.Int).SetString(in.Shares, 10)

	return MigrateL1Shares{
		GUID:              uuid.New(),
		UnstakerAddress:   common.HexToAddress(in.Withdrawer),
		StrategyAddress:   common.HexToAddress(in.Strategies),
		SharesAmount:      share,
		L1UnstakeMsgNonce: new(big.Int).SetUint64(in.L1UnStakeMessageNonce),
		ChainId:           ci,
		MTxHash:           common.Hash{},
		QTxHash:           common.Hash{},
		SourceHash:        sourceHash,
		Status:            PendingStatus,
		Timestamp:         time.Now().Unix(),
	}
}

func (m *migrateL1SharesDB) UpdateMigrateL1SharesTransactionHash(update MigrateL1Shares) error {
	result := m.gorm.Table("migrate_l1_shares").Where("guid = ?", update.GUID.String()).Updates(map[string]interface{}{"m_tx_hash": update.MTxHash.String()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *migrateL1SharesDB) UpdateQueueWithdrawalsTransactionHash(update MigrateL1Shares) error {
	result := m.gorm.Table("migrate_l1_shares").Where("guid = ?", update.GUID.String()).Updates(map[string]interface{}{"q_tx_hash": update.QTxHash.String()})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *migrateL1SharesDB) MigrateL1SharesByTxHash(txHash string) (*MigrateL1Shares, error) {
	var migrateL1Shares MigrateL1Shares
	result := m.gorm.Table("migrate_l1_shares").Where("m_tx_hash = ?", txHash).Take(&migrateL1Shares)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &migrateL1Shares, nil
}

func (m *migrateL1SharesDB) MigrateL1SharesBySourceHash(sourceHash string) (*MigrateL1Shares, error) {
	var migrateL1Shares MigrateL1Shares
	result := m.gorm.Table("migrate_l1_shares").Where("source_hash = ?", sourceHash).Take(&migrateL1Shares)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &migrateL1Shares, nil
}

func (m *migrateL1SharesDB) ChangeMigrateL1SharesSentStatusByTxHash(txHash string) error {
	var migrateL1Shares MigrateL1Shares
	result := m.gorm.Table("migrate_l1_shares").Where("m_tx_hash = ?", txHash).Take(&migrateL1Shares)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	migrateL1Shares.Status = mSentStatus
	err := m.gorm.Save(&migrateL1Shares).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *migrateL1SharesDB) ChangeQueueWithdrawalsSentStatusByTxHash(txHash string) error {
	var migrateL1Shares MigrateL1Shares
	result := m.gorm.Table("migrate_l1_shares").Where("q_tx_hash = ?", txHash).Take(&migrateL1Shares)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	migrateL1Shares.Status = qSentStatus
	err := m.gorm.Save(&migrateL1Shares).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *migrateL1SharesDB) StoreMigrateL1Shares(updates []MigrateL1Shares) error {
	result := m.gorm.CreateInBatches(&updates, len(updates))
	return result.Error
}

func (m *migrateL1SharesDB) OldestPendingNoSentTransaction() (*MigrateL1Shares, error) {
	var migrateL1Shares MigrateL1Shares
	result := m.gorm.Table("migrate_l1_shares").Where("status = ? and m_tx_hash = ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&migrateL1Shares)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &migrateL1Shares, nil
}

func (m *migrateL1SharesDB) OldestMPendingSentTransaction() (*MigrateL1Shares, error) {
	var migrateL1Shares MigrateL1Shares
	result := m.gorm.Table("migrate_l1_shares").Where("status = ? and m_tx_hash != ?", PendingStatus, common.Hash{}.String()).Order("timestamp asc").Find(&migrateL1Shares)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &migrateL1Shares, nil
}

func (m *migrateL1SharesDB) OldestQPendingSentTransaction() (*MigrateL1Shares, error) {
	var migrateL1Shares MigrateL1Shares
	result := m.gorm.Table("migrate_l1_shares").Where("status = ? and q_tx_hash != ?", mSentStatus, common.Hash{}.String()).Order("timestamp asc").Find(&migrateL1Shares)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &migrateL1Shares, nil
}
