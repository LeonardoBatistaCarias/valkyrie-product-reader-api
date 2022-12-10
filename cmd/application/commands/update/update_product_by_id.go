package update

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/domain/product"

	uuid "github.com/satori/go.uuid"
)

type UpdateProductByIDCommandHandler interface {
	Handle(ctx context.Context, cmd UpdateProductByIDCommand) error
}

type updateProductByIDHandler struct {
	gateway product.ProductGateway
}

func NewUpdateProductByIDHandler(productGateway product.ProductGateway) *updateProductByIDHandler {
	return &updateProductByIDHandler{gateway: productGateway}
}

func (c *updateProductByIDHandler) Handle(ctx context.Context, cmd UpdateProductByIDCommand) error {
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
	return c.gateway.UpdateProductByID(ctx, *product)
}
