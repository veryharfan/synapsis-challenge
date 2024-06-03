package entity

import "time"

type Transaction struct {
	Id         int64
	Invoice    string
	CustomerId int64
	ProductId  int64
	Quantity   int
	Amount     float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
