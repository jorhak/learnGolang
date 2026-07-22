package repository

import (
   "fmt"
)

type CreditCard struct {}

func (c CreditCard) Pay(amount float64) {
    fmt.Printf(
        "[CARD] %.2f\n",
         amount,
    )
}

type QRPayment struct {}

func (q QRPayment) Pay(amount float64) {
    fmt.Printf(
        "[QR] %.2f\n",
         amount,
    )
}
