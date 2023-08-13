package repository

import (
	"context"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/lucas-simao/kubernetes/api/internal/entity"
)

type Repository interface {
	PostCustomer(context.Context, entity.Customer) (entity.Customer, error)
	GetCustomers(context.Context) ([]entity.Customer, error)
	GetCustomerById(context.Context, uuid.UUID) (entity.Customer, error)
	PatchCustomerById(context.Context, entity.Customer) (entity.Customer, error)
	DeleteCustomerById(context.Context, uuid.UUID) error
	ApplyMigration(path string)
}
