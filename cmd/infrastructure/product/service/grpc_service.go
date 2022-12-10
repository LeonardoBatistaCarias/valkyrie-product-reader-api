package service

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/create"
	deleteCommand "github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/delete"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/update"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/queries/get_by"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/proto/pb"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/proto/pb/model"
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

func (s *grpcService) CreateProduct(ctx context.Context, req *pb.CreateProductReq) (*pb.CreateProductRes, error) {
	p := req.GetProduct()
	command := create.NewCreateProductCommand(p.GetProductID(), p.GetName(), p.GetDescription(), p.GetBrand(), float64(p.GetPrice()), p.GetQuantity(), uuid.FromStringOrNil(p.GetCategoryID()), nil, p.GetActive())

	if err := s.ps.Commands.CreateProduct.Handle(ctx, *command); err != nil {
		log.Printf("CreateProduct.Handle %s", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	return &pb.CreateProductRes{Product: p}, nil
}

func (s *grpcService) GetProductByID(ctx context.Context, req *pb.GetProductByIDReq) (*pb.GetProductByIDRes, error) {
	query := get_by.NewGetProductByIdQuery(req.ProductID)
	p, err := s.ps.Queries.GetProductById.Handle(ctx, query)
	if err != nil {
		log.Printf("GetProductByID.Handle %s", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}
	pp := &model.Product{ProductID: p.ProductID, Name: p.Name,
		Description:   p.Description,
		Brand:         int32(p.Brand),
		Price:         float32(p.Price),
		Quantity:      p.Quantity,
		CategoryID:    p.CategoryID.String(),
		ProductImages: nil,
		Active:        p.Active,
		CreatedAt:     p.CreatedAt.String(),
		UpdatedAt:     p.UpdatedAt.String(),
		DeletedAt:     p.DeletedAt.String()}

	res := &pb.GetProductByIDRes{Product: pp}
	return res, nil
}

func (s *grpcService) DeleteProductByID(ctx context.Context, req *pb.DeleteProductByIDReq) (*pb.DeleteProductByIDRes, error) {
	command := deleteCommand.NewDeleteProductByIDCommand(req.ProductID)

	if err := s.ps.Commands.DeleteProductByID.Handle(ctx, *command); err != nil {
		log.Printf("DeleteProductByID.Handle %s", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	return &pb.DeleteProductByIDRes{}, nil
}

func (s *grpcService) UpdateProductByID(ctx context.Context, req *pb.UpdateProductByIDReq) (*pb.UpdateProductByIDRes, error) {
	p := req.GetProduct()
	command := update.NewUpdateProductByIDCommand(p.GetProductID(), p.GetName(), p.GetDescription(), p.GetBrand(), float64(p.GetPrice()), p.GetQuantity(), uuid.FromStringOrNil(p.GetCategoryID()), nil, p.GetActive())

	if err := s.ps.Commands.UpdateProductByID.Handle(ctx, *command); err != nil {
		log.Printf("UpdateProductBytID.Handle %s", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	return &pb.UpdateProductByIDRes{}, nil
}

func (s *grpcService) errResponse(c codes.Code, err error) error {
	return status.Error(c, err.Error())
}
