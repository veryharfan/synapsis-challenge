package contract

import "time"

type CartRequest struct {
	CustomerId int64
	ProductId  int64 `json:"productId"`
	Quantity   int   `json:"quantity"`
}

type CartResponse struct {
	Id        int64           `json:"id"`
	Product   ProductResponse `json:"product"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
