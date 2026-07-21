package main

import (
      "fmt"
      "log"
      "erp-cli/internal/repository"
      "erp-cli/internal/service"
      "erp-cli/internal/domain"
)

func calculateTax(amount float64) float64{
    return amount * 0.13
}

func descount(total float64) float64{
    return total * 0.04
}

func main() {
    company := "Tech Corp"
    sales := 5000.0
    tax := calculateTax(sales)
    discount := descount(sales)
    total := (sales + tax) - discount

    fmt.Println("===== ERP REPORT =====")
    fmt.Println("Empresa:", company)
    fmt.Println("Ventas:", sales)
    fmt.Println("Impuestos:", tax)
    fmt.Println("Descuento:", discount) 
    fmt.Println("Total final:", total)
    fmt.Println("Estructura nueva")
    products := []domain.Product{
            {
               ID: 1,
               Name: "Yuta",
               Price: 83,
               Stock: 3,
            },
            {
               ID: 2,
               Name: "Maki",
               Price: 33,
               Stock: 33,
            },
            {
               ID: 1,
               Name: "Togen",
               Price: 13,
               Stock: 90,
            },
    }
    productService := service.NewProductService(
                  repository.JSONRepository{
                     FileName: "sal.json",
                  },
    )
    err := productService.AddProducts(products)
    if err != nil {
          log.Fatal("No se agregaron los productos", err)
    }
    product := domain.Product{
             ID: 99,
             Name: "jon snow",
             Price: 889,
             Stock: 9,
    }
    err = productService.AddProduct(product)
    if err != nil {
          log.Fatal("No se agrego el producto", err)
    }
    items, err := productService.GetProducts()
    if err != nil {
          log.Fatal("No se mostraron los productos", err)
    }
    for _, product := range items {
       fmt.Printf(
           "ID: %d | Name:%s | Price:%.2f | Stock:%d\n",
            product.ID,
            product.Name,
            product.Price,
            product.Stock,
       )
    }
}
