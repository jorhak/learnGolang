package main

import (
   "log"
   "erp-customer-storage/internal/domain"
   "erp-customer-storage/internal/repository"
   "erp-customer-storage/internal/service"
)

func main() {
    customers := []domain.Customer{
                {
                   ID: 1,
                   Name: "jon snow",
                   Email: "jon@snow.com",
                },
                {
                   ID: 2,
                   Name: "sansa stark",
                   Email: "sansa@stark.com",
                },
                {
                   ID: 3,
                   Name: "daneris targerian",
                   Email: "daneris@targerian.com",
                },
     }
    json := repository.NewJSONRepository("file.json")
    customerService := service.NewCustomerService(json)
    err := customerService.Store(customers)
    if err != nil {
      log.Fatal("No se almacenaron los clientes",err)
    }
    mysql := repository.NewMYSQLRepository("mysql.json")
    customerService = service.NewCustomerService(mysql)
    err = customerService.Store(customers)
    if err != nil {
      log.Fatal("No se almacenaron los clientes en MSYSQL", err)
    }
}
