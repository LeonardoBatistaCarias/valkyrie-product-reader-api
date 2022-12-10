package service

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/create"
	deleteByID "github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/delete"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/update"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/queries"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/queries/get_by"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/product"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/infrastructure/repository"
)

type ProductService struct {
	Commands *commands.ProductCommands
	Queries  *queries.ProductQueries
}

func NewProductService(
	mongoRepo repository.Repository,
) *ProductService {
	mongoDBGateway := product.NewProductMongoDBGateway(mongoRepo)

	createProductHandler := create.NewCreateProductHandler(mongoDBGateway)
	deleteProductByIDHandler := deleteByID.NewDeleteProductByIDHandler(mongoDBGateway)
	updateProductByIDHandler := update.NewUpdateProductByIDHandler(mongoDBGateway)

	getProductByIdHandler := get_by.NewGetProductByIdHandler(mongoDBGateway)

	productCommands := commands.NewProductCommands(createProductHandler, deleteProductByIDHandler, updateProductByIDHandler)
	productQueries := queries.NewProductQueries(getProductByIdHandler)

	return &ProductService{Commands: productCommands, Queries: productQueries}
}
