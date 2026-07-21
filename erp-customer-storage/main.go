package main

import (
   "fmt"
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
    jsonCustomer, err := customerService.Load()
    if err != nil {
      log.Fatal("No se pudo visualizar JSON",err)
    }
    for _, customer := range jsonCustomer {
       fmt.Printf(
           "ID: %d | Name: %s | Email: %s\n",
           customer.ID,
           customer.Name,
           customer.Email,
       )
    }
    mysql := repository.NewMYSQLRepository("mysql.json")
    customerService = service.NewCustomerService(mysql)
    err = customerService.Store(customers)
    if err != nil {
      log.Fatal("No se almacenaron los clientes en MSYSQL", err)
    }
    mysqlCustomer, err := customerService.Load()
    if err != nil {
      log.Fatal("No se pudo visualizar MYSQL",err)
    }
    for _, customer := range mysqlCustomer {
       fmt.Printf(
           "ID: %d :: Name: %s :: Email: %s\n",
           customer.ID,
           customer.Name,
           customer.Email,
       )
    }
}
