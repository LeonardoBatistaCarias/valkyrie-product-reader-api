package get_by

type GetProductByIdQuery struct {
	ProductID string
}

func NewGetProductByIdQuery(productID string) *GetProductByIdQuery {
	return &GetProductByIdQuery{ProductID: productID}
}
