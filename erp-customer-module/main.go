package main

import "fmt"

type Customer struct {
    ID int
    Name string
    Email string
}

func NewCustomer(
    id int,
    name string,
    email string,
) *Customer{
   return &Customer{
         ID: id,
         Name: name,
         Email: email,
   }
}

func (c *Customer) UpdateName(name string) {
    c.Name = name
}

func (c *Customer) UpdateEmail(email string) {
    c.Email = email
}

func (c *Customer) Display() {
    fmt.Printf(
        "ID: %d; Name: %s; Email:%s\n",
         c.ID,
         c.Name,
         c.Email,
    )
}

func main() {
    customer := NewCustomer(
             3,
             "Jon Snow",
             "jon@snow.com",
    )
    fmt.Println("Valor actual")
    customer.Display()
    customer.UpdateName("Jorhak Mormont")
    customer.UpdateEmail("jorhak@mormont.com")
    fmt.Println("Valores actuales")
    customer.Display()
}
