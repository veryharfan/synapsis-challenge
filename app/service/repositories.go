package service

import (
	"context"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/entity"
)

type CustomerRepository interface {
	GetByUsername(ctx context.Context, username string) (*entity.Customer, error)
	Create(ctx context.Context, customer entity.Customer) error
}

type ProductRepository interface {
	GetByCategory(ctx context.Context, param contract.GetListProductParam) ([]entity.Product, error)
	GetCountByCategory(ctx context.Context, param contract.GetListProductParam) (int, error)
	GetByIds(ctx context.Context, ids []int64) ([]entity.Product, error)
}

type CartRepository interface {
	Insert(ctx context.Context, param contract.CartRequest) error
	Update(ctx context.Context, param contract.CartRequest) error
	GetByCustomerIdAndProductId(ctx context.Context, customerId, productId int64) (*entity.Cart, error)
	GetByCustomerId(ctx context.Context, customerId int64) ([]entity.Cart, error)
}
