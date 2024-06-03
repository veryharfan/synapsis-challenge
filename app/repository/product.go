package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/entity"
	"synapsis-challenge/app/service"

	log "github.com/sirupsen/logrus"
)

const (
	AllFields = "id, name, category, price, stock, created_at, updated_at"
)

type productRepository struct {
	db *sql.DB
}

func InitProductRepository(db *sql.DB) service.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetByCategory(ctx context.Context, param contract.GetListProductParam) ([]entity.Product, error) {
	query := fmt.Sprintf("select %s from product", AllFields)
	args := []interface{}{}

	if param.Category != "" {
		query += " where category = $" + strconv.Itoa(len(args)+1)
		args = append(args, param.Category)
	}

	query += " order by id " + param.Sort + " limit $" + strconv.Itoa(len(args)+1) + " offset $" + strconv.Itoa(len(args)+2)
	args = append(args, param.Limit, param.Offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Errorf("getProductByCategory err:%v", err)
		return nil, err
	}

	defer rows.Close()

	products := []entity.Product{}
	for rows.Next() {
		p := entity.Product{}
		err = rows.Scan(&p.Id, &p.Name, &p.Category, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			log.Errorf("getProductByCategory err:%v", err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *productRepository) GetCountByCategory(ctx context.Context, param contract.GetListProductParam) (int, error) {
	query := "select count(1) from product"
	args := []interface{}{}

	if param.Category != "" {
		query += " where category = $1"
		args = append(args, param.Category)
	}

	var count int
	err := r.db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		log.Errorf("getProductByCategory err:%v", err)
		return 0, err
	}

	return count, nil
}
