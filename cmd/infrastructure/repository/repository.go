package repository

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/domain/product"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	CreateProduct(ctx context.Context, product *product.Product) (*product.Product, error)
	GetProductById(ctx context.Context, uuid uuid.UUID) (*product.Product, error)
}
