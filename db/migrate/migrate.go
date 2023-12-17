package migrate

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// performs database migrations
func MigrateDB(DB *sql.DB, dbName string) error {

	// Create a database connection instance
	instance, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		log.Println("Failed to create database connection instance:", err)
		return err
	}

	// Create a migration instance
	m, err := migrate.NewWithDatabaseInstance("file://db/migrate/migrations", dbName, instance)
	if err != nil {
		log.Println("Failed to create migration instance:", err)
		return err
	}

	// Execute migration
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Println("Failed to perform migration:", err)
		return err
	}

	log.Println("Database migration successful")
	return nil
}