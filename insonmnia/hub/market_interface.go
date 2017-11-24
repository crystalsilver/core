package hub

type Market interface {
	// OrderExists checks whether an order with the specified ID exists in the
	// marketplace.
	OrderExists(ID string) (bool, error)
	// CancelOrder removes order from marketplace
	CancelOrder(ID string) error
}
