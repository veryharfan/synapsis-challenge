package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/entity"
	"synapsis-challenge/app/service"

	log "github.com/sirupsen/logrus"
)

const (
	AllFieldsProduct = "id, name, category, price, stock, created_at, updated_at"
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
	query := fmt.Sprintf("select %s from product", AllFieldsProduct)
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

func (r *productRepository) GetByIds(ctx context.Context, ids []int64) ([]entity.Product, error) {
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = ids[i]
	}

	query := fmt.Sprintf("select %s from product where id in (%s)", AllFieldsProduct, strings.Join(placeholders, ", "))
	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var p entity.Product

		err := rows.Scan(&p.Id, &p.Name, &p.Category, &p.Price, &p.Stock, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
