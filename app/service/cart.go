package service

import (
	"context"
	"database/sql"
	"errors"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/entity"

	log "github.com/sirupsen/logrus"
)

type CartRepository interface {
	Insert(ctx context.Context, param contract.CartRequest) error
	Update(ctx context.Context, param contract.CartRequest) error
	GetByCustomerIdAndProductId(ctx context.Context, customerId, productId int64) (*entity.Cart, error)
}

type CartService interface {
	AddToCart(ctx context.Context, param contract.CartRequest) error
}

type cartService struct {
	cartRepo CartRepository
}

func InitCartService(cartRepo CartRepository) CartService {
	return &cartService{
		cartRepo: cartRepo,
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
