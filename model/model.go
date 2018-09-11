package model

// Order is the aggregate root for a purchase order.
type Order struct {
	ID    int
	Items []*LineItem
}

// LineItems is a value object representing a product and sale price for an
// order.
type LineItem struct {
	ProductID   int
	Description string
	Price       float64 // of course this should be some more precise type
}
