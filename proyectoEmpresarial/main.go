package main

import "fmt"

type Product struct {
    Name string
    Price float64
}

type Customer struct {
    ID    int
    Name  string
    Email string
}

type Order struct {
    ID       int
    Product  Product
    Quantity int
}

func (p Product) Tax() float64 {
    return p.Price * 0.13
}

func (p Product) FinalPrice() float64 {
    return p.Price + p.Tax()
}

func (c Customer) Display() {
    fmt.Println("ID:", c.ID)
    fmt.Println("Name:", c.Name)
    fmt.Println("Email:", c.Email)
}

func (o Order) Total() float64 {
    return o.Product.Price * float64(o.Quantity)
}

func updateProduct(product *Product){
    product.Price = 1500
}

func main(){
    laptop := Product{
      Name: "Laptop Dell",
      Price: 1000,
    }

    customer := Customer{
      ID:    1,
      Name:  "Jon Snow",
      Email: "jon@test.com",
    }

    order := Order{
      ID: 1,
      Product: Product{
        Name:  "Laptop",
        Price: 1000,
      },
      Quantity: 3,
    }

    fmt.Println(laptop.FinalPrice())
    customer.Display()
    fmt.Println(order.Total())

    product := Product{
        Name: "Laptop",
        Price: 1200,
    }
    fmt.Println("Valor actual: ", product.Price)
    updateProduct(&product)
    fmt.Println("Valor actualizado: ", product.Price)
}
