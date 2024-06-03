package service

import (
	"context"
	"database/sql"
	"errors"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/entity"

	log "github.com/sirupsen/logrus"
)

type CartService interface {
	AddToCart(ctx context.Context, param contract.CartRequest) error
	GetByCustomerId(ctx context.Context, customerId int64) ([]contract.CartResponse, error)
}

type cartService struct {
	cartRepo    CartRepository
	productRepo ProductRepository
}

func InitCartService(cartRepo CartRepository, productRepo ProductRepository) CartService {
	return &cartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *cartService) AddToCart(ctx context.Context, param contract.CartRequest) error {
	cart, err := s.cartRepo.GetByCustomerIdAndProductId(ctx, param.CustomerId, param.ProductId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("AddToCart err: %v", err)
		return err
	}

	if cart == nil {
		err = s.cartRepo.Insert(ctx, param)
	} else {
		err = s.cartRepo.Update(ctx, param)
	}
	if err != nil {
		log.Errorf("AddToCart err: %v", err)
		return err
	}

	return nil
}

func (s *cartService) GetByCustomerId(ctx context.Context, customerId int64) ([]contract.CartResponse, error) {
	cart, err := s.cartRepo.GetByCustomerId(ctx, customerId)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return []contract.CartResponse{}, nil
	} else if err != nil {
		log.Errorf("GetCartByCustomerId err: %v", err)
		return nil, err
	}

	var productIds []int64
	for _, c := range cart {
		productIds = append(productIds, c.ProductId)
	}

	products, err := s.productRepo.GetByIds(ctx, productIds)
	if err != nil {
		log.Errorf("GetCartByCustomerId err: %v", err)
		return nil, err
	}

	mapProductsById := map[int64]entity.Product{}
	for _, p := range products {
		mapProductsById[p.Id] = p
	}

	var resp []contract.CartResponse
	for _, c := range cart {
		p := mapProductsById[c.ProductId]
		resp = append(resp, contract.CartResponse{
			Id: c.Id,
			Product: contract.ProductResponse{
				Id:        p.Id,
				Name:      p.Name,
				Category:  p.Category,
				Price:     p.Price,
				Stock:     p.Stock,
				CreatedAt: p.CreatedAt,
				UpdatedAt: p.UpdatedAt,
			},
			Quantity:  c.Quantity,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}

	return resp, nil
}
