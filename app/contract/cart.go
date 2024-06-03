package contract

type CartRequest struct {
	CustomerId int64
	ProductId  int64 `json:"productId"`
	Quantity   int   `json:"quantity"`
}
