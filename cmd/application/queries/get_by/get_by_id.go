package get_by

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/domain/product"
	"github.com/opentracing/opentracing-go"
)

type GetProductByIdHandler interface {
	Handle(ctx context.Context, query *GetProductByIdQuery) (*product.Product, error)
}

type getProductByIdHandler struct {
	gateway product.ProductGateway
}

func NewGetProductByIdHandler(gateway product.ProductGateway) *getProductByIdHandler {
	return &getProductByIdHandler{gateway: gateway}
}

func (q *getProductByIdHandler) Handle(ctx context.Context, query *GetProductByIdQuery) (*product.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getProductByIdHandler.Handle")
	defer span.Finish()

	product, err := q.gateway.GetProductByID(ctx, query.ProductID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
