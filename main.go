package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/evm-layer2/selaginella/db"
)

// The Config structure is used to map the structure of the TOML configuration file
type Config struct {
	AppName  string `toml:"app_name"`
	Port     int    `toml:"port"`
	Database struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
	} `toml:"database"`
}

func main() {
	// Read configuration file contents
	configFile, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}

	// Parse TOML
	var config Config
	err = toml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("Error unmarshalling TOML:", err)
	}

	// Print configuration information
	log.Printf("App Name: %s\n", config.AppName)
	log.Printf("Port: %d\n", config.Port)
	log.Printf("Database Host: %s\n", config.Database.Host)
	log.Printf("Database Port: %d\n", config.Database.Port)
	log.Printf("Database Username: %s\n", config.Database.Username)
	log.Printf("Database Password: %s\n", config.Database.Password)
	log.Printf("Database Name: %s\n", config.Database.Name)

	// Initialize database connection
	db.InitDB(
		config.Database.Username,
		config.Database.Password,
		config.Database.Name,
		config.Database.Host,
		config.Database.Port,
	)

	// Handle termination signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		// Close database connection when signal is received
		db.CloseDB()
		os.Exit(1)
	}()

}
