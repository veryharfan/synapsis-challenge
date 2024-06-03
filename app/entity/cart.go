package entity

import "time"

type Cart struct {
	Id         int64
	CustomerId int64
	ProductId  int64
	Quantity   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
