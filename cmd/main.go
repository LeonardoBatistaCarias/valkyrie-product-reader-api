package main

import (
	"flag"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/server"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/utils/logger"
	"log"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewAppLogger(cfg.Logger)
	logger.InitLogger()

	s := server.NewServer(logger, cfg)
	logger.Fatal(s.Run())
}
