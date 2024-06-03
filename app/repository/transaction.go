package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"synapsis-challenge/app/entity"
	"synapsis-challenge/app/service"
)

const (
	AllFieldsTransaction = "id, invoice, customer_id, product_id, quantity, amount, created_at, updated_at"
)

type transactionRepo struct {
	db *sql.DB
}

func InitTransactionRepo(db *sql.DB) service.TransactionRepository {
	return &transactionRepo{
		db: db,
	}
}

func (r *transactionRepo) Create(ctx context.Context, transactions []entity.Transaction) error {
	if len(transactions) == 0 {
		return errors.New("no transactions to insert")
	}

	valueStrings := make([]string, 0, len(transactions))
	valueArgs := make([]interface{}, 0, len(transactions)*5)

	for i, transaction := range transactions {
		// Create placeholders for each field
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		valueArgs = append(valueArgs, transaction.Invoice)
		valueArgs = append(valueArgs, transaction.CustomerId)
		valueArgs = append(valueArgs, transaction.ProductId)
		valueArgs = append(valueArgs, transaction.Quantity)
		valueArgs = append(valueArgs, transaction.Amount)
	}

	query := fmt.Sprintf("INSERT INTO transaction (invoice, customer_id, product_id, quantity, amount) VALUES %s",
		strings.Join(valueStrings, ","))

	_, err := r.db.ExecContext(ctx, query, valueArgs...)
	if err != nil {
		return err
	}

	return nil
}
