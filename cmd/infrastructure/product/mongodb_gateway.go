package product

import (
	"context"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/domain/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/product/persistence"
)

type MongoDBGateway struct {
	mongoRepo persistence.Repository
}

func NewProductMongoDBGateway(mongoRepo persistence.Repository) *MongoDBGateway {
	return &MongoDBGateway{
		mongoRepo: mongoRepo,
	}
}

func (g *MongoDBGateway) CreateProduct(ctx context.Context, p product.Product) error {
	if err := g.mongoRepo.CreateProduct(ctx, &p); err != nil {
		return err
	}

	return nil
}

func (g *MongoDBGateway) GetProductByID(ctx context.Context, productID string) (*product.Product, error) {
	p, err := g.mongoRepo.GetProductById(ctx, productID)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (g *MongoDBGateway) DeleteProductByID(ctx context.Context, productID string) error {
	if err := g.mongoRepo.DeleteProductByID(ctx, productID); err != nil {
		return err
	}

	return nil
}

func (g *MongoDBGateway) DeactivateProductByID(ctx context.Context, productID string) error {
	if err := g.mongoRepo.DeactivateProductByID(ctx, productID); err != nil {
		return err
	}

	return nil
}

func (g *MongoDBGateway) UpdateProductByID(ctx context.Context, product product.Product) error {
	if err := g.mongoRepo.UpdateProductByID(ctx, &product); err != nil {
		return err
	}

	return nil
}
