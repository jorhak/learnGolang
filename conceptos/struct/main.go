package main

import "fmt"


func main() {

    type Product struct {
        Name  string
        Price float64
        Stock int
    }

    type Customer struct {
        ID int
        Name string
        Email string
    }

    customer := Customer{
       ID: 12,
       Name: "Maria Contreras",
       Email: "macontreras@gmail.com",
    }

    laptop := Product{
       Name: "Laptop Dell",
       Price: 1200,
       Stock: 10,
    }

    fmt.Println(laptop.Name)
    fmt.Println(laptop.Price)
    fmt.Println(laptop.Stock)
    fmt.Println(laptop)

    fmt.Println(customer)

    products := []Product{
       {
          Name: "Laptop",
          Price: 1200,
          Stock: 10,
       },
       {
          Name: "Mouse",
          Price: 20,
          Stock: 50,
       },
    }

    fmt.Println("\n====== Slice de Structs  =========")
    for _, product := range products {
       fmt.Println(product.Name)
       fmt.Println(product.Price)
       fmt.Println(product.Stock)
    }
}
