package main

import (
	"flag"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/server"
	"log"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(cfg)
	log.Fatal(s.Run())
}
