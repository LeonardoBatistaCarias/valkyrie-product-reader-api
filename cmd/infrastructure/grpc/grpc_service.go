package grpc

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/create"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/update"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/queries/get_by"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/config"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/grpc/pb/model"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/grpc/pb/reader_service"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/metrics"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/product/service"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/utils/logger"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/utils/time_handler"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcService struct {
	cfg     *config.Config
	ps      *service.ProductService
	log     logger.Logger
	metrics *metrics.Metrics
}

func NewReaderGrpcService(cfg *config.Config, ps *service.ProductService, log logger.Logger, metrics *metrics.Metrics) *grpcService {
	return &grpcService{cfg: cfg, ps: ps, log: log, metrics: metrics}
}

func (s *grpcService) CreateProduct(ctx context.Context, req *reader_service.CreateProductReq) (*reader_service.CreateProductRes, error) {
	s.metrics.CreateProductGrpcRequests.Inc()

	p := req.GetProduct()
	command := create.NewCreateProductCommand(p.GetProductID(), p.GetName(), p.GetDescription(), p.GetBrand(), float64(p.GetPrice()), p.GetQuantity(), uuid.FromStringOrNil(p.GetCategoryID()), nil, p.GetActive())

	if err := s.ps.Commands.CreateProduct.Handle(ctx, *command); err != nil {
		s.log.WarnMsg("CreateProduct.Handle", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	s.metrics.SuccessGrpcRequests.Inc()
	return &reader_service.CreateProductRes{Product: p}, nil
}

func (s *grpcService) GetProductByID(ctx context.Context, req *reader_service.GetProductByIDReq) (*reader_service.GetProductByIDRes, error) {
	s.metrics.GetProductByIdGrpcRequests.Inc()

	query := get_by.NewGetProductByIdQuery(req.ProductID)
	p, err := s.ps.Queries.GetProductById.Handle(ctx, query)
	if err != nil {
		s.log.WarnMsg("GetProductById.Handle", err)
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
		UpdatedAt:     time_handler.ConvertTimeToString(p.UpdatedAt),
		DeletedAt:     time_handler.ConvertTimeToString(p.DeletedAt)}

	s.metrics.SuccessGrpcRequests.Inc()
	return &reader_service.GetProductByIDRes{Product: pp}, nil
}

func (s *grpcService) DeleteProductByID(ctx context.Context, req *reader_service.DeleteProductByIDReq) (*reader_service.DeleteProductByIDRes, error) {
	s.metrics.DeleteProductGrpcRequests.Inc()

	if err := s.ps.Commands.DeleteProductByID.Handle(ctx, req.GetProductID()); err != nil {
		s.log.WarnMsg("DeleteProduct.Handle", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	s.metrics.SuccessGrpcRequests.Inc()
	return &reader_service.DeleteProductByIDRes{}, nil
}

func (s *grpcService) DeactivateProductByID(ctx context.Context, req *reader_service.DeactivateProductByIDReq) (*reader_service.DeactivateProductByIDRes, error) {
	s.metrics.DeactivateProductGrpcRequests.Inc()

	if err := s.ps.Commands.DeactivateProductByID.Handle(ctx, req.GetProductID()); err != nil {
		s.log.WarnMsg("DeactivateProductByID.Handle", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	s.metrics.SuccessGrpcRequests.Inc()
	return &reader_service.DeactivateProductByIDRes{}, nil
}

func (s *grpcService) UpdateProductByID(ctx context.Context, req *reader_service.UpdateProductByIDReq) (*reader_service.UpdateProductByIDRes, error) {
	s.metrics.UpdateProductGrpcRequests.Inc()
	p := req.GetProduct()
	command := update.NewUpdateProductByIDCommand(p.GetProductID(), p.GetName(), p.GetDescription(), p.GetBrand(), float64(p.GetPrice()), p.GetQuantity(), uuid.FromStringOrNil(p.GetCategoryID()), nil, p.GetActive())

	if err := s.ps.Commands.UpdateProductByID.Handle(ctx, *command); err != nil {
		s.log.WarnMsg("UpdateProductBytID.Handle", err)
		return nil, s.errResponse(codes.InvalidArgument, err)
	}

	s.metrics.SuccessGrpcRequests.Inc()
	return &reader_service.UpdateProductByIDRes{}, nil
}

func (s *grpcService) errResponse(c codes.Code, err error) error {
	return status.Error(c, err.Error())
}
