package entity

import "time"

type TransactionStatus struct {
	Id        int64
	Invoice   string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
