package main

import (
         "fmt"
         "encoding/json"
         "os"
)

type Product struct {
    ID int
    Name string
    Price float64
    Stock int
}

type ProductRepository interface {
    Save(products []Product) error
    Load() ([]Product, error)
}

type JSONRepository struct {
    FileName string
}

func (r JSONRepository) Save(
    products []Product,
) error {
  data, err := json.MarshalIndent(
              products,
              "",
              " ",
  )
  if err != nil {
        return err
  }
  return os.WriteFile(
        r.FileName,
        data,
        0644,
  )
}

func (r JSONRepository) Load() (
    []Product, error,
) {
  content, err := os.ReadFile(
              r.FileName,
  )
  if err != nil {
        return nil, err
  }
  var products []Product
  err = json.Unmarshal(
       content,
       &products,
  )
  return products, err
}

type ServiceFormat struct {
    productRepository ProductRepository
}

func NewServiceFormat(
     productRepository ProductRepository,
) *ServiceFormat {
   return &ServiceFormat{
         productRepository: productRepository,
   }
}

func (s ServiceFormat) Save(products []Product) {
    s.productRepository.Save(products)
}

func (s ServiceFormat) Load() {
    products, err := s.productRepository.Load()
    if err != nil {
      fmt.Println("Error Load::ServiceFormat: ", err)
    }
    for _, product := range products {
       fmt.Printf(
           "%d | %s | %.2f | %d\n",
           product.ID,
           product.Name,
           product.Price,
           product.Stock,
       )
    }
}

func main() {

    products := []Product{
            {ID: 1,
             Name: "Maria",
             Price: 33,
             Stock: 2,
            },
             {ID: 2,
             Name: "Joaquina",
             Price: 88,
             Stock: 34,
            },
             {ID: 3,
             Name: "Susana",
             Price: 88,
             Stock: 9,
            },
    }
    service := NewServiceFormat(
           JSONRepository{
              FileName: "hola.json",
           },
    )
    service.Save(products)
    fmt.Println("Se guardaron los productos en un fichero json")
    service = NewServiceFormat(
             JSONRepository{
                FileName: "quetal.json",
             },
    )
    service.Load()
}






