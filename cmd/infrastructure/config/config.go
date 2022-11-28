package config

import (
	"flag"
	"fmt"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/mongodb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/utils/constants"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Reader microservice config path")
}

type Config struct {
	ServiceName      string           `mapstructure:"serviceName"`
	GRPC             GRPC             `mapstructure:"grpc"`
	Mongo            *mongodb.Config  `mapstructure:"mongo"`
	MongoCollections MongoCollections `mapstructure:"mongoCollections"`
}

type GRPC struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

type MongoCollections struct {
	Products string `mapstructure:"products"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.CONFIG_PATH)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/%s", getwd, constants.BASE_CONFIG_PATH)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.DEFAULT_CONFIG_TYPE)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	grpcPort := os.Getenv(constants.GRPC_PORT)
	if grpcPort != "" {
		cfg.GRPC.Port = grpcPort
	}

	mongoURI := os.Getenv(constants.MONGO_DB_URI)
	if mongoURI != "" {
		cfg.Mongo.URI = mongoURI
	}

	return cfg, nil
}
