package service

import (
	"context"
	"database/sql"
	"errors"
	"synapsis-challenge/app/contract"
	"synapsis-challenge/app/entity"
	"synapsis-challenge/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type CustomerRepository interface {
	GetByUsername(ctx context.Context, username string) (*entity.Customer, error)
	Create(ctx context.Context, customer entity.Customer) error
}

type CustomerService interface {
	Create(ctx context.Context, request contract.CustomerRequest) error
	Login(ctx context.Context, request contract.CustomerRequest) (contract.LoginResponse, error)
}

type customerService struct {
	customerRepository CustomerRepository
	jwtConfig          config.JWT
}

func InitCustomerService(customerRepository CustomerRepository, jwtConfig config.JWT) CustomerService {
	return &customerService{
		customerRepository: customerRepository,
		jwtConfig:          jwtConfig,
	}
}

func (s *customerService) Create(ctx context.Context, request contract.CustomerRequest) error {
	customer, err := s.customerRepository.GetByUsername(ctx, request.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if customer != nil {
		return contract.ErrUsernameAlreadyExist
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return err
	}

	err = s.customerRepository.Create(ctx, entity.Customer{
		Username: request.Username,
		Password: string(password),
	})

	return err
}

func (s *customerService) Login(ctx context.Context, request contract.CustomerRequest) (contract.LoginResponse, error) {
	customer, err := s.customerRepository.GetByUsername(ctx, request.Username)
	if err != nil {
		return contract.LoginResponse{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password)); err != nil {
		return contract.LoginResponse{}, contract.ErrWrongPassword
	}

	claims := jwt.MapClaims{
		"iss": "synapsis-challenge",
		"sub": customer.Username,
		"uid": customer.Id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(s.jwtConfig.SigningKey))
	if err != nil {
		return contract.LoginResponse{}, err
	}

	return contract.LoginResponse{Token: token}, nil
}
