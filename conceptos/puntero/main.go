package main

import "fmt"

func updatePrice(price *float64) {
    *price = 200
}

func main() {
    age := 30
    pointer := &age

    fmt.Println("Valor:", age)
    fmt.Println("Direccion:", pointer)
    fmt.Println("Valor:", *pointer)

    *pointer = 50
    fmt.Println(age)

    price := 100.0
    fmt.Println("Valor actual:", price)
    updatePrice(&price)
    fmt.Println("Valor actualizado:", price)
}
