package mongodb

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/utils/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI      string `mapstructure:"uri"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
}

func NewMongoDBConn(ctx context.Context, cfg *Config) (*mongo.Client, error) {

	client, err := mongo.NewClient(
		options.Client().ApplyURI(cfg.URI).
			SetAuth(options.Credential{Username: cfg.User, Password: cfg.Password}).
			SetConnectTimeout(constants.CONNECT_TIMEOUT).
			SetMaxConnIdleTime(constants.MAX_CONN_IDLE_TIME).
			SetMinPoolSize(constants.MIN_POOL_SIZE).
			SetMaxPoolSize(constants.MAX_POOL_SIZE))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
