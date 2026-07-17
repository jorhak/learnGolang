package main

import "fmt"

func calculateTax(amount float64) float64{
    return amount * 0.13
}

func descount(total float64) float64{
    return total * 0.04
}

func main() {
    company := "Tech Corp"
    sales := 5000.0
    tax := calculateTax(sales)
    discount := descount(sales)
    total := (sales + tax) - discount

    fmt.Println("===== ERP REPORT =====")
    fmt.Println("Empresa:", company)
    fmt.Println("Ventas:", sales)
    fmt.Println("Impuestos:", tax)
    fmt.Println("Descuento:", discount) 
    fmt.Println("Total final:", total)
}
