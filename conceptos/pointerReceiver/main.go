package main

import "fmt"

type Customer struct {
    Name string
}

func (c *Customer) ChangeName(name string) {
    c.Name = name
}

func main() {
    customer := Customer{
            Name: "Jon Snow",
    }
    fmt.Println("Nombre actual: ", customer.Name)

    customer.ChangeName("Juan Perez")
    fmt.Println("Nombre actualizado: ", customer.Name)
}
