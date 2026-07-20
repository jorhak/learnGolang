package main

import "fmt"

type Product struct {
   ID int
   Name string
   Price float64
   Stock int
}

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

func NewProduct(
     id int,
     name string,
     price float64,
     stock int,
) *Product {
   return &Product{
       ID: id,
       Name: name,
       Price: price,
       Stock: stock,
   }
}

func (p *Product) UpdatePrice(
    price float64,
) {
    p.Price = price
}

func (p *Product) AddStock(
    quantity int,
) {
   p.Stock += quantity
}

func (p Product) InventoryValue() float64 {
    return p.Price * float64(p.Stock)
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

   fmt.Println("======= Construct Pattern=======")
   product := NewProduct(
          1,
          "Laptop",
          1200,
          10,
    )

    product.UpdatePrice(1500)
    product.AddStock(5)
    fmt.Println(
       "Stock:", product.Stock,
    )
    fmt.Println(
       "Inventory Value:",
       product.InventoryValue(),
    )
}
