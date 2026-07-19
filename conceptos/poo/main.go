package main

import "fmt"

type Product struct {
    Name string
    Price float64
}

func (p Product) CalculateTax() float64{
    return p.Price * 0.13
}

func (p Product) TotalPrice() float64 {
    return p.Price + p.CalculateTax()
}

func main(){

    laptop := Product{
      Name:  "Dell Latitude",
      Price: 1200,
    }

    fmt.Println(laptop.Name)
    fmt.Println(laptop.Price)
    tax := laptop.CalculateTax()
    fmt.Println("Impuesto:", tax)
    fmt.Println("Total:", laptop.TotalPrice())
}
