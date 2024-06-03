package repository

import (
	"context"
	"database/sql"
	"fmt"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/entity"
	"synapsis-challenge/app/service"

	log "github.com/sirupsen/logrus"
)

const (
	AllFieldsCart = "id, customer_id, product_id, quantity, created_at, updated_at"
)

type cartRepository struct {
	db *sql.DB
}

func InitCartRepository(db *sql.DB) service.CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r *cartRepository) Insert(ctx context.Context, param contract.CartRequest) error {
	query := "insert into cart (customer_id, product_id, quantity) values ($1, $2, $3)"
	_, err := r.db.Exec(query, param.CustomerId, param.ProductId, param.Quantity)
	if err != nil {
		log.Errorf("InsertCart err:%v", err)
		return err
	}

	return nil
}

func (r *cartRepository) Update(ctx context.Context, param contract.CartRequest) error {
	query := `update cart set (quantity, updated_at) = ($1, now()) 
		where customer_id = $2 and product_id = $3`
	_, err := r.db.Exec(query, param.Quantity, param.CustomerId, param.ProductId)
	if err != nil {
		log.Errorf("UpdateCart err:%v", err)
		return err
	}

	return nil
}

func (r *cartRepository) GetByCustomerIdAndProductId(ctx context.Context, customerId, productId int64) (*entity.Cart, error) {
	var cart entity.Cart

	query := fmt.Sprintf("select %s from cart where customer_id = $1 and product_id = $2", AllFieldsCart)
	err := r.db.QueryRow(query, customerId, productId).
		Scan(&cart.Id, &cart.CustomerId, &cart.ProductId, &cart.Quantity, &cart.CreatedAt, &cart.UpdatedAt)
	if err != nil {
		log.Errorf("GetCartByCustomerIdAndProductId err:%v", err)
		return nil, err
	}

	return &cart, nil
}
