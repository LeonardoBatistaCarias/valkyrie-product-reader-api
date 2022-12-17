package product

import (
	"context"
)

type ProductGateway interface {
	CreateProduct(ctx context.Context, product Product) error
	GetProductByID(ctx context.Context, productID string) (*Product, error)
	DeleteProductByID(ctx context.Context, productID string) error
	DeactivateProductByID(ctx context.Context, productID string) error
	UpdateProductByID(ctx context.Context, product Product) error
}
