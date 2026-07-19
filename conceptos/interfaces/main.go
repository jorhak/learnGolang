package main

import "fmt"

type PaymentMethod interface {
    Pay(amount float64)
}

type CreditCard struct {
    Number string
}

type QRPayment struct {
    Code string
}

func (c CreditCard) Pay(amount float64) {
    fmt.Printf(
        "Pago con tarjeta: %.2f\n",
         amount,
    )
}

func (q QRPayment) Pay(amount float64) {
    fmt.Printf(
        "Pago QR: %.2f\n",
        amount,
    )
}

func ProcessPayment(
     method PaymentMethod,
     amount float64,
     ) {
         method.Pay(amount)
     }

func main() {
    card := CreditCard{
       Number: "123456",
    }

    qr := QRPayment{
       Code: "QR001",
    }

    ProcessPayment(card,100)
    ProcessPayment(qr, 200)
}
