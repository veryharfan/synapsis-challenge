package entity

import "time"

type Product struct {
	Id        int64
	Name      string
	Category  string
	Price     float64
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
