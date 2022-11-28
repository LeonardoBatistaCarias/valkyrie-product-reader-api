package queries

import (
	"github.com/LeonardoBatistaCarias/valkyrie-product-reader-api/cmd/application/queries/get_by"
)

type ProductQueries struct {
	GetProductById get_by.GetProductByIdHandler
}

func NewProductQueries(getProductById get_by.GetProductByIdHandler) *ProductQueries {
	return &ProductQueries{GetProductById: getProductById}
}
