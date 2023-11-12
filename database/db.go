// Database module defines the data DB struct which wraps specific DB interfaces for L1/L2 block headers, contract events, bridging schemas.
package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/evm-layer2/selaginella/config"
	_ "github.com/evm-layer2/selaginella/database/utils/serializers"
	"github.com/evm-layer2/selaginella/synchronizer/retry"
)

type DB struct {
	gorm   *gorm.DB
	Blocks AddressDB
}

func NewDB(dbConfig config.DBConfig) (*DB, error) {
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}

	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Name)
	if dbConfig.Port != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.Port)
	}
	if dbConfig.User != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.User)
	}
	if dbConfig.Password != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.Password)
	}

	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
	}

	gorm, err := retry.Do[*gorm.DB](context.Background(), 10, retryStrategy, func() (*gorm.DB, error) {
		gorm, err := gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gorm, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after multiple retries: %w", err)
	}

	db := &DB{
		gorm:   gorm,
		Blocks: NewAddressDB(gorm),
	}

	return db, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(tx *gorm.DB) error {
		return fn(dbFromGormTx(tx))
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func dbFromGormTx(tx *gorm.DB) *DB {
	return &DB{
		gorm:   tx,
		Blocks: NewAddressDB(tx),
	}
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

func NewRedis(redisCfg config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", redisCfg.Host, redisCfg.Port),
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err == nil && pong == "PONG" {
		return client, nil
	}
	return nil, err
}
