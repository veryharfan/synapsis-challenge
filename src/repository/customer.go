package repository

import (
	"context"
	"database/sql"
	"fmt"
	"synapsis-challenge/src/entity"
	"synapsis-challenge/src/service"
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
		QueryRow("select id, username, password from customer where username = $1", username).
		Scan(&customer.Id, &customer.Username, &customer.Password)
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
