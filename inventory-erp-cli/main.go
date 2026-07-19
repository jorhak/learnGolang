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

type Product struct {
   Name string
   Price float64
   Stock int
}

func printProducts(products []Product) {
    for _, product := range products {
       fmt.Printf(
           "%s | %.2f | %d\n",
           product.Name,
           product.Price,
           product.Stock,
       )
    }
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

    products := []Product{
        {
           Name: "Laptop",
           Price: 1200,
           Stock: 10,
        },
        {
           Name: "Mouse",
           Price: 20,
           Stock: 50,
        },
        {
           Name: "Keyboard",
           Price: 35,
           Stock: 25,
        },
    }
    fmt.Println("====== Slice Struct=======")
    printProducts(products)
}
