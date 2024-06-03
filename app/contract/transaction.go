package contract

type CheckoutRequest struct {
	CustomerId int64
	ProductIds []int64 `json:"productIds"`
}

type CheckoutResponse struct {
	PaymentUrl string `json:"paymentUrl"`
}
