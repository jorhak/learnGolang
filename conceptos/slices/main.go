package main

import "fmt"

func main(){

    products := []string{
        "Laptop",
        "Mouse",
        "Keyboard",
        "Monitor",
        "Printer",
    }
    fmt.Println(products)

    others := []string{}
    others = append(others, "Teclado")
    others = append(others, "Camara")
    others = append(others, "Microfono")

    fmt.Println(others)

    fmt.Println("Longitud de productos:", len(products))
    fmt.Println("Longitud de otros:", len(others))

    categories := []string{
        "Electronics",
        "Computers",
        "Accessories",
    }
    // Ignorara indice
    for _, categori := range categories {
        fmt.Println(categori)
    }

    for index, product := range products {
        fmt.Println(index,product)
    }
}
