package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/evm-layer2/selaginella/config"
	"github.com/evm-layer2/selaginella/deposit"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"strconv"
)

func main() {

	//Initialize log
	h := log.StreamHandler(os.Stdout, log.TerminalFormat(true))
	log.Root().SetHandler(h)

	//Read configuration file path from command line parameters
	//Configuration file path is config.toml by default
	var confPath = flag.String("c", "config.toml", "config path")
	flag.Parse()
	log.Info("config path", "path", *confPath)

	//Initialize configuration
	conf, err := config.New(*confPath)
	if err != nil {
		panic(err)
	}
	log.Info("config", "conf", conf)

	// Build database connection string
	dataSourceString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name)
	log.Info("database connection string", "dataSourceString", dataSourceString)

	// Open database connection
	db, err := sql.Open("postgres", dataSourceString)
	if err != nil {
		log.Error("failed to open database connection: %s", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Error("failed to close database connection: %s", err)
		}
	}(db)
	log.Info("database connection opened")

	// Check database connection
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	err = db.Ping()
	if err != nil {
		log.Error("failed to ping database: %s", err)
	}
	log.Info("database connection pinged")

	// Perform database migration
	instance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Error("failed to create database connection instance, err: ", err)
	}
	log.Info("database connection instance created")

	databaseInstance, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		conf.Database.Name,
		instance,
	)
	if err != nil {
		log.Error("Failed to create databaseInstance: ", err)
		panic(err)
	}
	log.Info("database instance created")

	if err = databaseInstance.Up(); !errors.Is(err, migrate.ErrNoChange) && err != nil {
		log.Error("failed to perform migration, err: ", err)
		panic(err)
	}
	log.Info("database migration performed")

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()
	log.Info("grpc server initialized")

	// Initialize deposit service
	depositService := &deposit.DepositService{Db: db}
	deposit.RegisterDepositServiceServer(grpcServer, depositService)
	log.Info("deposit service initialized")

	// Register gRPC service
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(conf.Rpc.Port))
	if err != nil {
		log.Error("net listen failed", "err", err)
		panic(err)
	}
	log.Info("net listen success", "port", conf.Rpc.Port)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Error("grpc server serve failed", "err", err)
		panic(err)
	}
	log.Info("grpc server serve success")

}
