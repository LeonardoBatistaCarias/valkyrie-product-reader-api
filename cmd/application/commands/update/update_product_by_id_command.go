package update

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type UpdateProductByIDCommand struct {
	ProductID     string
	Name          string
	Description   string
	Brand         int32
	Price         float64
	Quantity      int32
	CategoryID    uuid.UUID
	ProductImages []*UpdateProductByIDImageCommand
	Active        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type UpdateProductByIDImageCommand struct {
	Address string
}

func NewUpdateProductByIDCommand(
	productID string,
	name string,
	description string,
	brand int32,
	price float64,
	quantity int32,
	categoryID uuid.UUID,
	images []*UpdateProductByIDImageCommand,
	active bool) *UpdateProductByIDCommand {
	return &UpdateProductByIDCommand{
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
