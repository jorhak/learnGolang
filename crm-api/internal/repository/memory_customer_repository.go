package repository

import (
	"crm-api/internal/domain"
	"errors"
)

// MemoryCustomerRepository implementa CustomerRepository utilizando memoria
type MemoryCustomerRepository struct {
	customers []domain.Customer
}

// NewMemoryCustomerRepository crea un nuevo repositorio en memoria
func NewMemoryCustomerRepository() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: []domain.Customer{
			{
				ID: 1,
				Name: "Maria Joaquina",
				Email: "maria@joaquina.com",
			},
		},
	}
}

// GetAll devuelve todos los clientes alamacenados en memoria
func (r *MemoryCustomerRepository) GetAll() []domain.Customer {
	return r.customers
}

// GetByID devuelve al cliente con el ID correspondiente
func (r *MemoryCustomerRepository) GetByID(id int) (domain.Customer, error) {
	for _, customer := range r.customers  {
		if customer.ID == id {
			return customer, nil
		}
	}
	return domain.Customer{}, errors.New("Customer not found")
}

// Crear un customer en memoria
func (r *MemoryCustomerRepository) Create(customer domain.Customer) error {
	r.customers = append(r.customers, customer)
	return nil
}

// Elimina un customer de memoria
func (r *MemoryCustomerRepository) Delete(id int) error {
	r.customers = eliminarByID(r.customers, id)
	return nil
}

func eliminarByID(customers []domain.Customer, id int) []domain.Customer {
	for i := range customers {
		if customers[i].ID == id {
			return append(customers[:i],customers[i+1:]...)
		}
	}
	return customers
}
