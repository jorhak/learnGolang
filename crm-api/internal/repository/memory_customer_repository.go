package repository

import (
	"crm-api/internal/domain"
)

// MemoryCustomerRepository implementa CustomerRepository utilizando memoria
type MemoryCustomerRepository struct {}

// NewMemoryCustomerRepository crea un nuevo repositorio en memoria
func NewMemoryCustomerRepository() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{}
}

// GetAll devuelve todos los clientes alamacenados en memoria
func (r *MemoryCustomerRepository) GetAll() []domain.Customer {
	return []domain.Customer{
		{
			ID: 1,
			Name: "Glucosa",
			Email: "glucosa@gmail.com",
		},
		{
			ID: 2,
			Name: "Creatina",
			Email: "creatina@gmail.com",
		},
	}
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
