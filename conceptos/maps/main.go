package main

import "fmt"

func main(){

    inventory := make(map[string]int)

    inventory["Laptop"] = 15
    inventory["Mouse"] = 50
    inventory["Keyboard"] = 20

    fmt.Println(inventory)
    fmt.Println(inventory["Laptop"])

    for producto, stock := range inventory{

        fmt.Println("Producto: ", producto, "  ","Stock", stock)
    }
}
