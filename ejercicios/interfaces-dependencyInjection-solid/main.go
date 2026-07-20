package main

import "fmt"

type PaymentMethod interface {
    Pay(amount float64)
}

type CreditCard struct {
    Tipo string
}
// En las struct podria dejarlo vacio
// type CreditCard struct {}
// Y solo lo tendira que lo llamo sin parametros
// card := CreditCard{}
// Esto para todos los struct
func (c CreditCard) Pay(amount float64){
    fmt.Printf(
        "Pago Tarjeta: %.2f - Tipo: %s\n",
         amount,
         c.Tipo,
    )
}

type QRPayment struct {
    Image string
}

func (q QRPayment) Pay(amount float64) {
    fmt.Printf(
        "Pago QR: %.2f - Imangen: %s\n",
        amount,
        q.Image,
    )
}

type BankTransfer struct {
    SRC string
    DEST string
}

func (b BankTransfer) Pay(amount float64) {
    fmt.Printf(
        "Pago Transferencia: %.2f - Fuente: %s - Destino: %s\n",
        amount,
        b.SRC,
        b.DEST,
    )
}

func ProcessPayment(
    method PaymentMethod,
    amount float64,
){
    method.Pay(amount)
}
// Dependency Injection
type PaymentService struct {
    method PaymentMethod
}
// Constructores
func NewPaymentService(
    method PaymentMethod,
) *PaymentService {
   return &PaymentService{
         method: method, 
  }
}
// Procesar
func (p PaymentService) Process(
    amount float64,
){
    p.method.Pay(amount)
}

func main() {
    card := CreditCard{
            Tipo: "Debito",
    }

    qr := QRPayment{
          Image: "QR893",
    }

    transfer := BankTransfer{
                SRC: "Banco Union",
                DEST: "Mercantil Santa Cruz",
    }

    ProcessPayment(card, 100)
    ProcessPayment(qr, 50)
    ProcessPayment(transfer, 200)
    fmt.Println("======== Dependency Injection ============")
    service := NewPaymentService(
              QRPayment{
                  Image: "PA839",
              },
    )
    service.Process(40)
    service = NewPaymentService(
              CreditCard{
                   Tipo: "Master Card",
              },
    )
    service.Process(90)
    service = NewPaymentService(
             BankTransfer{
                  SRC: "Banco Fie",
                  DEST: "Banco Economico",
             },
    )
    service.Process(44)
}
