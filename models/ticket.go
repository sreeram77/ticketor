package models

type Ticket struct {
	ID      string
	UserID  User
	From    string
	To      string
	Number  string
	Section string
	Price   Money
}

type Money struct {
	Amount   float64
	Currency string
}
