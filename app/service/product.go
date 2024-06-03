package service

import (
	"context"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/entity"

	"github.com/sirupsen/logrus"
)

type ProductRepository interface {
	GetByCategory(ctx context.Context, param contract.GetListProductParam) ([]entity.Product, error)
	GetCountByCategory(ctx context.Context, param contract.GetListProductParam) (int, error)
}

type ProductService interface {
	GetByCategory(ctx context.Context, param contract.GetListProductParam) ([]contract.ProductResponse, *contract.Pagination, error)
}

type productService struct {
	productRepo ProductRepository
}

func InitProductService(productRepo ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) GetByCategory(ctx context.Context, param contract.GetListProductParam) ([]contract.ProductResponse, *contract.Pagination, error) {
	products, err := s.productRepo.GetByCategory(ctx, param)
	if err != nil {
		logrus.Error("GetByCategory ProductService err:", err)
		return nil, nil, err
	}

	productCount, err := s.productRepo.GetCountByCategory(ctx, param)
	if err != nil {
		logrus.Error("GetCountByCategory ProductService err:", err)
		return nil, nil, err
	}

	var resp []contract.ProductResponse
	for _, p := range products {
		resp = append(resp, contract.ProductResponse{
			Id:        p.Id,
			Name:      p.Name,
			Category:  p.Category,
			Price:     p.Price,
			Stock:     p.Stock,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	pagination := contract.GetPaginationData(param.Page, param.Limit, productCount)

	return resp, pagination, nil
}
