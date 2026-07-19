package main

import "fmt"

type Product struct {
    ID       int
    Name     string
    Price    float64
    Stock    int
    Category string
}

func addProduct(product Product, products []Product) []Product{
    return append(products, product)
}

func printProducts(products []Product) {
    for _, product := range products {
       fmt.Printf(
          "%d %s %.2f %d %s\n",
          product.ID,
          product.Name,
          product.Price,
          product.Stock,
          product.Category,
       )
    }
}

func calculateInventoryValue(products []Product) float64 {
    inventoryValue := 0.0
    for _, product := range products {
       inventoryValue += product.Price * float64(product.Stock)
    }
    return inventoryValue
}

func main() {

   laptop := Product{
      ID:       1,
      Name:     "Laptop",
      Price:    1200,
      Stock:    10,
      Category: "Electronics",
   }

   mouse := Product{
     ID:       2,
     Name:     "Mouse",
     Price:    20,
     Stock:    50,
     Category: "Accessories",
   }

   products := []Product{}
   products = addProduct(laptop,products)
   products = addProduct(mouse,products)
   printProducts(products)
   fmt.Printf("\nInventory Value: %.2f\n", calculateInventoryValue(products))

}
