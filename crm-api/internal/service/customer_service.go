package service

import(
   "crm-api/internal/domain"
	 "crm-api/internal/repository"
)

// CustomerService contiene la logica del negocio 
type CustomerService struct {
	repository repository.CustomerRepository
}

// NewCustomerService crea un nuevo servicio
func NewCustomerService(
	repository repository.CustomerRepository,
) *CustomerService {
    return &CustomerService{
			repository: repository,
		}
}

// GetCustomers devuelve todos los clientes
func (s *CustomerService) GetCustomers() []domain.Customer {
    return s.repository.GetAll()
}

// Funcion que busca el objeto con el ID
func (s *CustomerService) GetCustomerByID(id int) (domain.Customer, bool) {
	return s.repository.GetByID(id)
}
