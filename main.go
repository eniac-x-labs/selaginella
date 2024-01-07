package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/log"
	"github.com/evm-layer2/selaginella/config"
)

func main() {
	var f = flag.String("c", "config.yml", "config path")
	flag.Parse()
	conf, err := config.New(*f)
	if err != nil {
		panic(err)
	}
	log.Info("conf", conf)
}
