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

type CryptoPayment struct {
    Wallet string
}

type BankTransfer struct {
    Src  string
    Dest string
    Mount float64
}

type PaypalPayment struct {
    Dest string
    Mount float64
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

func (c CryptoPayment) Pay(amount float64) {
    fmt.Printf(
        "Billetera: %s monto %.2f\n",
         c.Wallet,
         amount,
    )
}

func (b BankTransfer) Pay(amount float64) {
    fmt.Printf(
        "Origen: %s Destino: %s Monto: %.2f Ingreso: %.2f\n",
         b.Src, b.Dest, b.Mount, amount,
    )
}

func (p PaypalPayment) Pay(amount float64) {
    fmt.Printf(
        "Destino: %s Monto: %.2f Ingreso: %.2f\n",
         p.Dest, p.Mount, amount,
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

    crypto := CryptoPayment{
       Wallet: "Bitcoin",
    }

    banck := BankTransfer{
       Src: "100002423",
       Dest: "84794",
       Mount: 67.00,
    }

    p := PaypalPayment{
      Dest: "8889274",
      Mount: 89.99,
    }

    ProcessPayment(card,100)
    ProcessPayment(qr, 200)
    ProcessPayment(crypto,400)
    ProcessPayment(banck,900)
    ProcessPayment(p,90)
}
