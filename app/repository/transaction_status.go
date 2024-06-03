package repository

import (
	"context"
	"database/sql"
	"synapsis-challenge/app/entity"
	"synapsis-challenge/app/service"

	log "github.com/sirupsen/logrus"
)

type transactionStatusRepo struct {
	db *sql.DB
}

func InitTransactionStatusRepo(db *sql.DB) service.TransactionStatusRepository {
	return &transactionStatusRepo{
		db: db,
	}
}

func (r *transactionStatusRepo) Create(ctx context.Context, transactionStatus entity.TransactionStatus) error {
	query := "INSERT INTO transaction_status (invoice, status) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, transactionStatus.Invoice, transactionStatus.Status)
	if err != nil {
		log.Error("CreateTransactionStatus err:", err)
		return err
	}

	return nil
}

func (r *transactionStatusRepo) UpdateByInvoice(ctx context.Context, invoice, status string) error {
	query := "update transaction_status set status = $1 where invoice = $2"
	_, err := r.db.ExecContext(ctx, query, status, invoice)
	if err != nil {
		log.Error("UpdateTransactionStatusByInvoice err:", err)
		return err
	}

	return nil
}
