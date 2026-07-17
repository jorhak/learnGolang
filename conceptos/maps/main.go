package main

import "fmt"

func printInventory(inventory map[string]int){
    for producto, stock := range inventory{
       fmt.Println(producto, stock)
    }
}

func main(){

    inventory := make(map[string]int)

    inventory["Laptop"] = 15
    inventory["Mouse"] = 50
    inventory["Keyboard"] = 20

    printInventory(inventory)
}
