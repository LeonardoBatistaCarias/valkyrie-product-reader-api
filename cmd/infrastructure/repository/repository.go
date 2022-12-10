package repository

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/domain/product"
)

type Repository interface {
	CreateProduct(ctx context.Context, product *product.Product) (*product.Product, error)
	GetProductById(ctx context.Context, productID string) (*product.Product, error)
}
