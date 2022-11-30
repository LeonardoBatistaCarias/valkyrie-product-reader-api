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

func (g *ProductMongoDBGateway) Create(ctx context.Context, p product.Product) error {
	_, err := g.mongoRepo.CreateProduct(ctx, &p)
	if err != nil {
		return err
	}

	return nil
}
