package repository

import (
	"crm-api/internal/domain"
)

// CustomerRepository define el contrato para acceder a los clientes
type CustomerRepository interface {
	GetAll() []domain.Customer
	GetByID(id int) (domain.Customer, bool)
}
