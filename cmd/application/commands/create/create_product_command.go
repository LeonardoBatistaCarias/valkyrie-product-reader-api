package create

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreateProductCommand struct {
	ProductID     uuid.UUID
	Name          string
	Description   string
	Brand         int32
	Price         float64
	Quantity      int32
	CategoryID    uuid.UUID
	ProductImages []*CreateProductImageCommand
	Active        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type CreateProductImageCommand struct {
	Address string
}

func NewCreateProductCommand(
	productID uuid.UUID,
	name string,
	description string,
	brand int32,
	price float64,
	quantity int32,
	categoryID uuid.UUID,
	images []*CreateProductImageCommand,
	active bool) *CreateProductCommand {
	return &CreateProductCommand{
		ProductID:     productID,
		Name:          name,
		Description:   description,
		Brand:         brand,
		Price:         price,
		Quantity:      quantity,
		CategoryID:    categoryID,
		ProductImages: images,
		Active:        active,
	}
}
