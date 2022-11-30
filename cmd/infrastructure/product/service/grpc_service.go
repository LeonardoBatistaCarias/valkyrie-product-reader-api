package service

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/create"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/config"
	protoProduct "github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/proto/product"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type grpcService struct {
	cfg *config.Config
	ps  *ProductService
}

func NewReaderGrpcService(cfg *config.Config, ps *ProductService) *grpcService {
	return &grpcService{cfg: cfg, ps: ps}
}

func (s *grpcService) CreateProduct(ctx context.Context, req *protoProduct.Product) (*protoProduct.Product, error) {
	command := create.NewCreateProductCommand(uuid.FromStringOrNil(req.GetProductID()), req.GetName(), req.GetDescription(), 1, float64(req.GetPrice()), 1, uuid.NewV4(), nil, true)

	if err := s.ps.Commands.CreateProduct.Handle(ctx, *command); err != nil {
		log.Printf("CreateProduct.Handle %s", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	return req, nil
}

func (s *grpcService) errResponse(c codes.Code, err error) error {
	return status.Error(c, err.Error())
}
