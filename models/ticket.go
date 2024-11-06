package models

import "fmt"

type Ticket struct {
	ID      string
	UserID  string
	From    string
	To      string
	Number  string
	Section string
	Price   Money
	User    User
}

type Money struct {
	amount   float64
	currency string
	symbol   string
}

func NewMoney(amount float64) Money {
	return Money{
		amount:   amount,
		currency: "USD",
		symbol:   "$",
	}
}

func (m Money) String() string {
	return fmt.Sprintf("%f %s", m.amount, m.symbol)
}
