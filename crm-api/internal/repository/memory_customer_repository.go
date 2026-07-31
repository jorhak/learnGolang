package repository

import (
	"crm-api/internal/domain"
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
func (r *MemoryCustomerRepository) GetByID(id int) (domain.Customer, bool) {
	customers := r.GetAll()
	customer := bucarByID(id, customers)
	if customer != nil {
		return *customer, true
	}
	return domain.Customer{}, false
}

func bucarByID(id int, customers []domain.Customer) *domain.Customer {
	for i := range customers {
		if customers[i].ID == id {
			return &customers[i]
		}
	}
	return nil
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
