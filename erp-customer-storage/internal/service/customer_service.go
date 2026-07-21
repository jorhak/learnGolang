package service

import (
  "erp-customer-storage/internal/domain"
)

type CustomerService struct {
    customerRepository domain.CustomerRepository
}

func NewCustomerService(
    customerRepository domain.CustomerRepository,
) *CustomerService {
  return &CustomerService{
        customerRepository: customerRepository,
  }
}

func (s *CustomerService) Store(customers []domain.Customer) error {
    return s.customerRepository.Store(customers)
}

