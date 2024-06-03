package repository

import (
	"context"
	"database/sql"
	"fmt"
	"synapsis-challenge/app/entity"
	"synapsis-challenge/app/service"
)

type customerRepository struct {
	db *sql.DB
}

func InitCustomerRepository(db *sql.DB) service.CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) GetByUsername(ctx context.Context, username string) (*entity.Customer, error) {
	customer := entity.Customer{}
	err := r.db.
		QueryRow("select id, username, password, created_at, updated_at from customer where username = $1", username).
		Scan(&customer.Id, &customer.Username, &customer.Password, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &customer, nil

}

func (r *customerRepository) Create(ctx context.Context, customer entity.Customer) error {
	_, err := r.db.Exec(`
		INSERT INTO customer (username, password) 
		VALUES ($1, $2)
	`, customer.Username, customer.Password)

	return err
}
