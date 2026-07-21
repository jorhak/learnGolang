package service

import (
  "erp-cli/internal/domain"
)

type ProductService struct {
    productRepository domain.ProductRepository
}

func NewProductService(
    productRepository domain.ProductRepository,
) *ProductService {
   return &ProductService{
         productRepository: productRepository,
   }
}

func (s *ProductService) AddProducts(products []domain.Product) error {
    return s.productRepository.Save(products)
}

func (s *ProductService) AddProduct(product domain.Product) error {
    products, err := s.productRepository.Load()
    if err != nil {
          return err
    }
    products = append(products, product)
    return s.productRepository.Save(products)
}

func (s *ProductService) GetProducts() ([]domain.Product, error) {
    return s.productRepository.Load()
}
