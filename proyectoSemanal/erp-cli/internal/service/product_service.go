package service

import (
   "proyectoSemanal/erp-cli/internal/domain"
)

type ProductService struct {
    Repository domain.ProductRepository
}

func NewProductService(
    repository domain.ProductRepository,
) *ProductService {
  return &ProductService{
     Repository: repository,
  }
}

func (s *ProductService) AddProducts(products []domain.Product) error {
    return s.Repository.Save(products)
}

func (s *ProductService) GetProducts() ([]domain.Product, error) {
    return s.Repository.Load()
}
