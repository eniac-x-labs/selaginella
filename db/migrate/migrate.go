package migrate

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

// performs database migrations
func MigrateDB(DB *sql.DB, dbName string) error {
	// Create a migration instance
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		return err
	}

	// Create migration
	m, err := migrate.NewWithDatabaseInstance("file://./db/migrate/migrations", dbName, driver)
	if err != nil {
		return err
	}

	// Execute migration
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	fmt.Println("Database migration successful")
	return nil
}