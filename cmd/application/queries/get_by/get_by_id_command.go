package get_by

import uuid "github.com/satori/go.uuid"

type GetProductByIdQuery struct {
	ProductID uuid.UUID
}

func NewGetProductByIdQuery(productID uuid.UUID) *GetProductByIdQuery {
	return &GetProductByIdQuery{ProductID: productID}
}
