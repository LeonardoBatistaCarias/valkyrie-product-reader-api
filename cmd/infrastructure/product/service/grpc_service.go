package service

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/create"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/queries/get_by"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/config"
	readerService "github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/proto/product_reader"
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

func (s *grpcService) CreateProduct(ctx context.Context, req *readerService.CreateProductReq) (*readerService.CreateProductRes, error) {
	command := create.NewCreateProductCommand(uuid.FromStringOrNil(req.GetProductID()), req.GetName(), req.GetDescription(), 1, req.GetPrice(), 1, uuid.NewV4(), nil, true)

	if err := s.ps.Commands.CreateProduct.Handle(ctx, *command); err != nil {
		log.Printf("CreateProduct.Handle %s", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	return &readerService.CreateProductRes{ProductID: req.GetProductID()}, nil
}

func (s *grpcService) GetProductById(ctx context.Context, req *readerService.GetProductByIdReq) (*readerService.GetProductByIdRes, error) {
	productUUID, err := uuid.FromString(req.GetProductID())
	if err != nil {
		log.Printf("uuid.FromString %s", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	query := get_by.NewGetProductByIdQuery(productUUID)
	product, err := s.ps.Queries.GetProductById.Handle(ctx, query)
	print(product)
	if err != nil {
		log.Fatalf("GetProductById.Handle %s", err)
		return nil, s.errResponse(codes.Internal, err)
	}

	return nil, nil
	//return &readerService.GetProductByIdRes{Product: models.ProductToGrpcMessage(product)}, nil
}

func (s *grpcService) errResponse(c codes.Code, err error) error {
	return status.Error(c, err.Error())
}
