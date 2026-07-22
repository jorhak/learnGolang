package main

import(
   "fmt"
   "proyectoSemanal/erp-cli/internal/service"
   "proyectoSemanal/erp-cli/internal/domain"
   "proyectoSemanal/erp-cli/internal/repository"
)

func main() {
    products := []domain.Product{
            {
               ID: 1,
               Name: "Laptop",
               Price: 1200,
               Stock: 10,
            },
            {
               ID: 2,
               Name: "Mouse",
               Price: 20,
               Stock: 50,
            },
            {
               ID: 3,
               Name: "Monitor",
               Price: 300,
               Stock: 15,
            },
    }

    customer := domain.Customer{
                ID: 1,
                Name: "Jon Snow",
                Email: "jon@snow.com",
    }

    order := domain.Order{
             ID: 1,
             Product: products[0],
             Customer: customer,
             Quantity: 2,
    }
    json := repository.NewJSONRepository("data/products.json")
    productService := service.NewProductService(json)
    productService.AddProducts(products)
    loadedProducts, err := productService.GetProducts()

    if err != nil {
      panic(err)
    }

    fmt.Printf(
        "Productos cargados: %d\n\n",
         len(loadedProducts),
    )
    fmt.Printf(
        "Cliente:\n%s\n\n",
        customer.Name,
    )
    fmt.Println("Orden creada\n")
    total := order.Total()
    fmt.Printf(
        "Total:\n%.2f\n\n",
         total,
    )

    paymentService := service.NewPaymentService(repository.CreditCard{},)
    paymentService.Process(total)
}
