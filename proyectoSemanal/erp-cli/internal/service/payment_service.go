package service

import (
   "proyectoSemanal/erp-cli/internal/domain"
)

type PaymentService struct {
    method domain.PaymentMethod
}

func NewPaymentService(
         method domain.PaymentMethod,
) *PaymentService {
    return &PaymentService{
               method: method,
    }
}

func (s PaymentService) Process(amount float64,) {
    s.method.Pay(amount)
}
