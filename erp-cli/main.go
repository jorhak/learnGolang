package main

import "fmt"

func calculateTax(amount float64) float64{
    return amount * 0.13
}

func main() {
    company := "ACME ERP"
    sales := 1500.0
    tax := calculateTax(sales)
    total := sales + tax

    fmt.Println("Empresa:", company)
    fmt.Println("Ventas:", sales)
    fmt.Println("Impuesto:", tax)
    fmt.Println("Total:", total)
}
