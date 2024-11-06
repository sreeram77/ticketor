package models

type Ticket struct {
	ID     string
	UserID string
	From   string
	To     string
	Number string
	Price  Money
}

type Money struct {
	Amount   float64
	Currency string
}
