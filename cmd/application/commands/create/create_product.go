package create

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/domain/product"

	uuid "github.com/satori/go.uuid"
)

type CreateProductCommandHandler interface {
	Handle(ctx context.Context, cmd CreateProductCommand) error
}

type createProductHandler struct {
	gateway product.ProductGateway
}

func NewCreateProductHandler(productGateway product.ProductGateway) *createProductHandler {
	return &createProductHandler{gateway: productGateway}
}

func (c *createProductHandler) Handle(ctx context.Context, cmd CreateProductCommand) error {
	product := product.NewProduct(
		cmd.ProductID,
		cmd.Name,
		cmd.Description,
		1,
		cmd.Price,
		cmd.Quantity,
		uuid.NewV4(),
		nil,
		true)
	return c.gateway.CreateProduct(ctx, *product)
}
