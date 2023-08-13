package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lucas-simao/kubernetes/api/internal/entity"
)

func (r *repository) PostCustomer(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	namedStmt, err := r.db.PrepareNamedContext(ctx, sqlCreateCustomer)
	if err != nil {
		return entity.Customer{}, entity.PostCustomerError
	}

	var newCustomer entity.Customer

	err = namedStmt.GetContext(ctx, &newCustomer, customer)
	if err != nil {
		return entity.Customer{}, entity.PostCustomerError
	}

	return newCustomer, nil
}

func (r *repository) GetCustomers(ctx context.Context) ([]entity.Customer, error) {
	rows, err := r.db.QueryContext(ctx, sqlGetCustomers)
	if err != nil {
		return nil, err
	}

	var customers []entity.Customer

	for rows.Next() {
		var c entity.Customer

		err := rows.Scan(
			&c.Id,
			&c.FirstName,
			&c.LastName,
			&c.BirthDate,
			&c.UpdatedAt,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (r *repository) GetCustomerById(ctx context.Context, customerId uuid.UUID) (entity.Customer, error) {
	row := r.db.QueryRowContext(ctx, sqlGetCustomerById, customerId)

	var customer entity.Customer

	err := row.Scan(
		&customer.Id,
		&customer.FirstName,
		&customer.LastName,
		&customer.BirthDate,
		&customer.UpdatedAt,
		&customer.CreatedAt,
	)
	if err != nil {
		return entity.Customer{}, entity.GetCustomerError
	}

	return customer, nil
}

func (r *repository) PatchCustomerById(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	namedStmt, err := r.db.PrepareNamedContext(ctx, sqlUpdateCustomerById)
	if err != nil {
		return entity.Customer{}, entity.PutCustomerError
	}

	var customerUpdated entity.Customer

	err = namedStmt.GetContext(ctx, &customerUpdated, customer)
	if err != nil {
		return entity.Customer{}, entity.PutCustomerError
	}

	return customerUpdated, nil
}

func (r *repository) DeleteCustomerById(ctx context.Context, customerId uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, sqlDeleteCustomerById, customerId)

	if err != nil {
		return entity.DeleteCustomerError
	}

	return nil
}
