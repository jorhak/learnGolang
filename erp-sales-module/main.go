package main

import "fmt"

type Product struct {
    ID    int
    Name  string
    Price float64
}

type Customer struct {
    ID   int
    Name string
}

type Order struct {
    ID       int
    Product  Product
    Customer Customer
    Quantity int
}

func (p Product) Tax() float64 {
    return p.Price * 0.13
}

func (c Customer) Display() {
    fmt.Printf(
        "Customer: %s\n",
         c.Name,
    )
}

func (o Order) Total() float64 {
    return o.Product.Price * float64(o.Quantity)
}

func main(){
    product := Product{
        ID:   1,
        Name: "Laptop",
        Price: 1200,
    }

    customer := Customer{
        ID:   1,
        Name: "Jon Snow",
    }

    order := Order{
        ID:       1,
        Product:  product,
        Customer: customer,
        Quantity: 2,
    }

    customer.Display()
    fmt.Println(
             "Order Total:",
              order.Total(),
    )

}
