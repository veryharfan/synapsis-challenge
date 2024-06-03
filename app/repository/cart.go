package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
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

func (r *cartRepository) GetByCustomerId(ctx context.Context, customerId int64) ([]entity.Cart, error) {
	query := fmt.Sprintf("select %s from cart where customer_id = $1", AllFieldsCart)
	rows, err := r.db.Query(query, customerId)
	if err != nil {
		return nil, err
	}

	var cart []entity.Cart
	for rows.Next() {
		var item entity.Cart
		err = rows.Scan(&item.Id, &item.CustomerId, &item.ProductId, &item.Quantity, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			log.Errorf("GetCartByCustomerId err:%v", err)
			return nil, err
		}
		cart = append(cart, item)
	}

	return cart, nil
}

func (r *cartRepository) Delete(ctx context.Context, customerId, productId int64) error {
	query := "delete from cart where customer_id = $1 and product_id = $2"
	_, err := r.db.Exec(query, customerId, productId)
	if err != nil {
		log.Fatal(err)
	}

	return nil

}

func (r *cartRepository) GetByCustomerIdAndProductIds(ctx context.Context, customerId int64, productIds []int64) ([]entity.Cart, error) {
	placeholders := make([]string, len(productIds))
	args := make([]interface{}, len(productIds)+1)
	args[0] = customerId

	for i := range productIds {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
		args[i+1] = productIds[i]
	}

	query := fmt.Sprintf("select %s from cart where customer_id = $1 and product_id in (%s)", AllFieldsCart, strings.Join(placeholders, ", "))
	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Error("GetByCustomerIdAndProductIds err:", err)
		return nil, err
	}
	defer rows.Close()

	var cart []entity.Cart

	for rows.Next() {
		var c entity.Cart

		err := rows.Scan(&c.Id, &c.CustomerId, &c.ProductId, &c.Quantity, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		cart = append(cart, c)
	}

	return cart, nil
}

func (r *cartRepository) DeleteByCustomerIdAndProductIds(ctx context.Context, customerId int64, productIds []int64) error {
	placeholders := make([]string, len(productIds))
	args := make([]interface{}, len(productIds)+1)
	args[0] = customerId

	for i := range productIds {
		placeholders[i] = fmt.Sprintf("$%d", i+2)
		args[i+1] = productIds[i]
	}

	query := fmt.Sprintf("delete from cart where customer_id = $1 and product_id in (%s)", strings.Join(placeholders, ", "))
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Error("DeleteCartByCustomerIdAndProductIds err:", err)
		return err
	}

	return nil
}
