package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/evm-layer2/selaginella/common/retry"
	"github.com/evm-layer2/selaginella/config"
	_ "github.com/evm-layer2/selaginella/database/utils/serializers"
)

type DB struct {
	gorm                             *gorm.DB
	CrossChainTransfer               CrossChainTransferDB
	UpdateDepositFundingPoolBalance  UpdateDepositFundingPoolBalanceDB
	UpdateWithdrawFundingPoolBalance UpdateWithdrawFundingPoolBalanceDB
	UnstakeBatch                     UnstakeBatchDB
	MigrateL1Shares                  MigrateL1SharesDB
	TransferToL2Bridge               TransferToL2BridgeDB
	BatchMint                        BatchMintDB
	TransferL2Share                  TransferL2ShareDB
}

func NewDB(ctx context.Context, dbConfig config.DB) (*DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.DbHost, dbConfig.DbName)
	if dbConfig.DbPort != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.DbPort)
	}
	if dbConfig.DbUser != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.DbUser)
	}
	if dbConfig.DbPassword != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.DbPassword)
	}
	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
	}
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	gorm, err := retry.Do[*gorm.DB](context.Background(), 10, retryStrategy, func() (*gorm.DB, error) {
		gorm, err := gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gorm, nil
	})
	if err != nil {
		return nil, err
	}
	db := &DB{
		gorm:                             gorm,
		CrossChainTransfer:               NewCrossChainTransferDB(gorm),
		UpdateDepositFundingPoolBalance:  NewDepositUpdateFundingPoolBalanceDB(gorm),
		UpdateWithdrawFundingPoolBalance: NewWithdrawUpdateFundingPoolBalanceDB(gorm),
		UnstakeBatch:                     NewUnstakeBatchDB(gorm),
		MigrateL1Shares:                  NewMigrateL1SharesDB(gorm),
		TransferToL2Bridge:               NewTransferToL2BridgeDB(gorm),
		BatchMint:                        NewBatchMintDB(gorm),
		TransferL2Share:                  NewTransferL2ShareDB(gorm),
	}
	return db, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(tx *gorm.DB) error {
		txDB := &DB{
			gorm:                             tx,
			CrossChainTransfer:               NewCrossChainTransferDB(tx),
			UpdateDepositFundingPoolBalance:  NewDepositUpdateFundingPoolBalanceDB(tx),
			UpdateWithdrawFundingPoolBalance: NewWithdrawUpdateFundingPoolBalanceDB(tx),
			UnstakeBatch:                     NewUnstakeBatchDB(tx),
			MigrateL1Shares:                  NewMigrateL1SharesDB(tx),
			TransferToL2Bridge:               NewTransferToL2BridgeDB(tx),
			BatchMint:                        NewBatchMintDB(tx),
			TransferL2Share:                  NewTransferL2ShareDB(tx),
		}
		return fn(txDB)
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}

		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}
		return nil
	})
	return err
}
