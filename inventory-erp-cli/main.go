package main

import "fmt"

func printInventory(inventory map[string]int){
    fmt.Println("\n====== INVENTORY =========\n")

    for product, stock := range inventory{
       fmt.Printf("%s -> %d\n", product, stock)
    }
}

func calculateTotalStock(inventory map[string]int) int{
    total := 0

    for _, stock := range inventory{
        total+= stock
    }
    return total
}

func main(){
    inventory := map[string]int{
        "Laptop":   15,
        "Mouse":    50,
        "Keyboard": 30,
        "Monitor":  10,
    }

    printInventory(inventory)
    total := calculateTotalStock(inventory)
    fmt.Println("\nTotal Stock:", total)
}
