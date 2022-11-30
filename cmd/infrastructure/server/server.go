package server

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/mongodb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/product/service"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/repository"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	cfg         *config.Config
	mongoClient *mongo.Client
	ps          *service.ProductService
}

func NewServer(cfg *config.Config) *server {
	return &server{cfg: cfg}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	mongoDBConn, err := mongodb.NewMongoDBConn(ctx, s.cfg.Mongo)
	if err != nil {
		return errors.Wrap(err, "NewMongoDBConn")
	}
	s.mongoClient = mongoDBConn
	defer mongoDBConn.Disconnect(ctx) // nolint: errcheck
	log.Printf("Mongo connected: %v", mongoDBConn.NumberSessionsInProgress())

	mongoRepo := repository.NewMongoRepository(s.cfg, s.mongoClient)

	s.ps = service.NewProductService(mongoRepo)

	closeGrpcServer, grpcServer, err := s.NewGrpcServer()
	if err != nil {
		return errors.Wrap(err, "NewScmGrpcServer")
	}
	defer closeGrpcServer() // nolint: errcheck

	<-ctx.Done()
	grpcServer.GracefulStop()
	return nil
}
