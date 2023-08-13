package customers

import (
	"context"

	"github.com/google/uuid"
	"github.com/lucas-simao/kubernetes/api/internal/entity"
)

type Service interface {
	PostCustomer(context.Context, entity.Customer) (entity.Customer, error)
	GetCustomers(context.Context) ([]entity.Customer, error)
	GetCustomerById(context.Context, uuid.UUID) (entity.Customer, error)
	PatchCustomerById(context.Context, entity.Customer) (entity.Customer, error)
	DeleteCustomerById(context.Context, uuid.UUID) error
}
