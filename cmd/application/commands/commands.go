package commands

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/create"
	deleteByID "github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/delete"
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/update"
)

type ProductCommands struct {
	CreateProduct         create.CreateProductCommandHandler
	DeleteProductByID     deleteByID.DeleteProductByIDCommandHandler
	DeactivateProductByID deleteByID.DeactivateProductByIDCommandHandler
	UpdateProductByID     update.UpdateProductByIDCommandHandler
}

func NewProductCommands(
	createProduct create.CreateProductCommandHandler,
	deleteProductByID deleteByID.DeleteProductByIDCommandHandler,
	deactivateProductByID deleteByID.DeactivateProductByIDCommandHandler,
	updateProductByID update.UpdateProductByIDCommandHandler,
) *ProductCommands {
	return &ProductCommands{
		CreateProduct:         createProduct,
		DeleteProductByID:     deleteProductByID,
		DeactivateProductByID: deactivateProductByID,
		UpdateProductByID:     updateProductByID,
	}
}
