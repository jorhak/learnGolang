package main

import "fmt"

type PaymentMethod interface {
    Pay(amount float64)
}

type CreditCard struct {
    NroCuenta string
}

type QRPayment struct {
    SRC   string
    DEST  string
    MONTO float64
}

type BankTransfer struct {
    BANK_SRC  string
    BANK_DEST string
    MONTO     float64
}

func (c CreditCard) Pay(amount float64) {
    fmt.Printf(
        "Pago Tarjeta: %.2f\n",
        amount,
    )
}


func (qr QRPayment) Pay(amount float64) {
    fmt.Printf(
        "Pago QR: %.2f\n",
        amount,
    )
}


func (b BankTransfer) Pay(amount float64) {
    fmt.Printf(
        "Pago Transferencia: %.2f\n",
         amount,
    )
}

func ProcessPayment(method PaymentMethod, amount float64) {
    method.Pay(amount)
}

// Dependency Injection
type PaymentService struct {
    method PaymentMethod
}

func NewPaymentService(
    method PaymentMethod,
) *PaymentService {
  return &PaymentService{
        method: method,
  }
}

func (p PaymentService) Process(
    amount float64,
) {
  p.method.Pay(amount)
}

func main(){
    credit := CreditCard{
          NroCuenta: "1002743",
    }

    qr := QRPayment{
          SRC: "984723",
          DEST: "1002874",
          MONTO: 84.00,
    }

    b := BankTransfer{
         BANK_SRC: "Banco Union",
         BANK_DEST: "Banco Bisa",
         MONTO: 890.00,
    }
    ProcessPayment(credit, 100)
    ProcessPayment(qr, 200)
    ProcessPayment(b, 120)
    fmt.Println("====== Dependency Injection ========")
    cardService := NewPaymentService(
                      credit,
    )
     cardService.Process(39)
     qrService := NewPaymentService(
                      qr,
    )
    qrService.Process(15)
    transferService := NewPaymentService(
                      b,
    )
    transferService.Process(990)
}
