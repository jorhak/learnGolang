package main

import "fmt"

type CustomerRepository interface {
    Save(customer *Customer)
}

type PostgresRepository struct {}

func (p PostgresRepository) Save(
     customer *Customer,
) {
    fmt.Println(
        "Guardado en PostgreSQL",
    )
}

type MongoRepository struct {}

func (m MongoRepository) Save(
     customer *Customer,
) {
    fmt.Println(
        "Guardado en MongoDB",
    )
}

type CustomerService struct {
    repository CustomerRepository
}

func NewCustomerService(
    repository CustomerRepository,
) *CustomerService {
    return &CustomerService{
          repository: repository,
    }
}

func (c CustomerService) Process(customer *Customer){
    c.repository.Save(customer)
}

type Customer struct {
    ID int
    Name string
    Email string
}

func NewCustomer(
    id int,
    name string,
    email string,
) *Customer {
   return &Customer{
         ID: id,
         Name: name,
         Email: email,
   }
}

func main() {

    service := NewCustomerService(
              PostgresRepository{},
    )
    customer := NewCustomer(
                44,
                "Joaquina",
                "mjoaquina@gmail.com",
    )
    service.Process(customer)
    service = NewCustomerService(
             MongoRepository{},
    )
    service.Process(customer)
}
