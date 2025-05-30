package service

import (
	"aryavidyananta/Golang-Project/domain"
	"aryavidyananta/Golang-Project/dto"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type CustomerService struct {
	CustomerRepository domain.CustomerRepository
}

func NewCustomer(CustomerRepository domain.CustomerRepository) domain.CustomerService {
	return &CustomerService{
		CustomerRepository: CustomerRepository,
	}
}

// Index implements domain.CustomerService.
func (c *CustomerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.CustomerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var CustomerData []dto.CustomerData
	for _, v := range customers {
		CustomerData = append(CustomerData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}
	return CustomerData, nil
}

func (c CustomerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:        uuid.NewString(),
		Code:      req.Code,
		Name:      req.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	_, err := c.CustomerRepository.Save(ctx, &customer)
	return err
}

func (c CustomerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persited, err := c.CustomerRepository.FindById(ctx, req.ID)
	if err != nil {
		return err
	}
	if persited.ID == "" {
		return errors.New("customer not found")
	}
	persited.Code = req.Code
	persited.Name = req.Code
	persited.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	_, err = c.CustomerRepository.Update(ctx, &persited)
	return err
}

func (c CustomerService) Delete(ctx context.Context, id string) error {
	exist, err := c.CustomerRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if exist.ID == "" {
		return errors.New("customer not found")
	}
	return c.CustomerRepository.Delete(ctx, id)

}

func (c CustomerService) Show(ctx context.Context, id string) (dto.CustomerData, error) {
	persisted, err := c.CustomerRepository.FindById(ctx, id)
	if err != nil {
		return dto.CustomerData{}, err
	}
	if persisted.ID == "" {
		return dto.CustomerData{}, errors.New("customer not found")

	}
	return dto.CustomerData{
		ID:   persisted.ID,
		Code: persisted.Code,
		Name: persisted.Name,
	}, nil
}
