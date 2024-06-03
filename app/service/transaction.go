package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/entity"
	"time"

	log "github.com/sirupsen/logrus"
)

type TransactionService interface {
	Checkout(ctx context.Context, param contract.CheckoutRequest) (*contract.CheckoutResponse, error)
	Update(ctx context.Context, param contract.CallbackUpdateTransaction) error
}

type transactionService struct {
	db                    *sql.DB
	cartRepo              CartRepository
	productRepo           ProductRepository
	transactionRepo       TransactionRepository
	transactionStatusRepo TransactionStatusRepository
}

func InitTransactionService(db *sql.DB, cartRepo CartRepository, productRepo ProductRepository, transactionRepo TransactionRepository, transactionStatusRepo TransactionStatusRepository) TransactionService {
	return &transactionService{
		db:                    db,
		cartRepo:              cartRepo,
		productRepo:           productRepo,
		transactionRepo:       transactionRepo,
		transactionStatusRepo: transactionStatusRepo,
	}
}

func (s *transactionService) Checkout(ctx context.Context, param contract.CheckoutRequest) (*contract.CheckoutResponse, error) {
	carts, err := s.cartRepo.GetByCustomerIdAndProductIds(ctx, param.CustomerId, param.ProductIds)
	if err != nil {
		log.Error("Checkout err:", err)
		return nil, err
	}

	products, err := s.productRepo.GetByIds(ctx, param.ProductIds)
	if err != nil {
		log.Error("Checkout err:", err)
		return nil, err
	}

	mapProductsById := map[int64]entity.Product{}
	for _, p := range products {
		mapProductsById[p.Id] = p
	}

	invoice := "INV/" + strconv.Itoa(int(param.CustomerId)) + "/" + time.Now().Format("20060102150405")

	var total float64
	var transactions []entity.Transaction
	for _, c := range carts {
		product := mapProductsById[c.ProductId]
		amount := float64(c.Quantity) * product.Price

		transactions = append(transactions, entity.Transaction{
			Invoice:    invoice,
			CustomerId: param.CustomerId,
			ProductId:  c.ProductId,
			Quantity:   c.Quantity,
			Amount:     amount,
		})
		total += amount
	}

	err = s.cartRepo.DeleteByCustomerIdAndProductIds(ctx, param.CustomerId, param.ProductIds)
	if err != nil {
		log.Error("Checkout err:", err)
		return nil, err
	}

	err = s.transactionRepo.Create(ctx, transactions)
	if err != nil {
		log.Error("Checkout err:", err)
		return nil, err
	}

	err = s.transactionStatusRepo.Create(ctx, entity.TransactionStatus{
		Invoice: invoice,
		Status:  "pending",
	})
	if err != nil {
		log.Error("Checkout err:", err)
		return nil, err
	}

	paymentUrl := fmt.Sprintf("payment-url/invoice=%s&amount=%s", invoice, strconv.Itoa(int(total)))

	return &contract.CheckoutResponse{
		PaymentUrl: paymentUrl,
	}, nil
}

func (s *transactionService) Update(ctx context.Context, param contract.CallbackUpdateTransaction) error {
	err := s.transactionStatusRepo.UpdateByInvoice(ctx, param.Invoice, param.Status)
	if err != nil {
		log.Error("Update transaction status err:", err)
		return err
	}

	return nil
}
