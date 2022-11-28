package commands

import "github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/commands/create"

type ProductCommands struct {
	CreateProduct create.CreateProductCommandHandler
}

func NewProductCommands(createProduct create.CreateProductCommandHandler) *ProductCommands {
	return &ProductCommands{CreateProduct: createProduct}
}
