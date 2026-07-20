package main

import "fmt"

type Product struct {
    ID int
    Name string
    Price float64
}

func NewProduct(
    id int,
    name string,
    price float64,
) *Product {
  return &Product{
        ID: id,
        Name: name,
        Price: price,
  }
}

func (p Product) Display() {
    fmt.Printf(
        "ID: %d, Name: %s, Price: %.2f\n",
        p.ID, p.Name, p.Price,
    )
}

func main(){
    product := NewProduct(
           1,
           "Laptop",
           1200,
    )

    product.Display()
}
