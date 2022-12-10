package product

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/repository"
)

type ProductMongoDBGateway struct {
	mongoRepo repository.Repository
}

func NewProductMongoDBGateway(mongoRepo repository.Repository) *ProductMongoDBGateway {
	return &ProductMongoDBGateway{
		mongoRepo: mongoRepo,
	}
}

func (g *ProductMongoDBGateway) CreateProduct(ctx context.Context, p product.Product) error {
	if err := g.mongoRepo.CreateProduct(ctx, &p); err != nil {
		return err
	}

	return nil
}

func (g *ProductMongoDBGateway) GetProductByID(ctx context.Context, productID string) (*product.Product, error) {
	p, err := g.mongoRepo.GetProductById(ctx, productID)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (g *ProductMongoDBGateway) DeleteProductByID(ctx context.Context, productID string) error {
	if err := g.mongoRepo.DeleteProductByID(ctx, productID); err != nil {
		return err
	}

	return nil
}

func (g *ProductMongoDBGateway) UpdateProductByID(ctx context.Context, product product.Product) error {
	if err := g.mongoRepo.UpdateProductByID(ctx, &product); err != nil {
		return err
	}

	return nil
}
