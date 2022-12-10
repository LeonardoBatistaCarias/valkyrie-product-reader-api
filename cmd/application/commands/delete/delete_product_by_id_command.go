package delete

type DeleteProductByIDCommand struct {
	ProductID string
}

func NewDeleteProductByIDCommand(
	productID string) *DeleteProductByIDCommand {
	return &DeleteProductByIDCommand{
		ProductID: productID,
	}
}
