package customers

import (
	"context"

	"github.com/google/uuid"
	"github.com/lucas-simao/kubernetes/api/internal/entity"
	"github.com/lucas-simao/kubernetes/api/internal/repository"
)

type service struct {
	repository repository.Repository
}

func New(r repository.Repository) Service {
	return service{
		repository: r,
	}
}

func (s service) PostCustomer(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	return s.repository.PostCustomer(ctx, customer)
}

func (s service) GetCustomers(ctx context.Context) ([]entity.Customer, error) {
	return s.repository.GetCustomers(ctx)
}

func (s service) GetCustomerById(ctx context.Context, customerId uuid.UUID) (entity.Customer, error) {
	return s.repository.GetCustomerById(ctx, customerId)
}

func (s service) PatchCustomerById(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	return s.repository.PatchCustomerById(ctx, customer)
}

func (s service) DeleteCustomerById(ctx context.Context, customerId uuid.UUID) (err error) {
	return s.repository.DeleteCustomerById(ctx, customerId)
}
