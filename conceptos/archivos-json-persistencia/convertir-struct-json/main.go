package main

import "fmt"
import "encoding/json"
import "os"

type Product struct {
    ID    int      `json:"id"`
    Name  string   `json:"name"`
    Price float64  `json:"price"`
    Stock int      `json:"stock"`
}

func main(){
    product := Product{
              ID:   1,
              Name: "Laptop",
              Price: 1200,
              Stock: 10,
    }

    data, err := json.MarshalIndent(
                product,
                "",
                " ",
    )
    if err != nil {
          panic(err)
    }
    fmt.Println(string(data))
    fmt.Println("Slice a json")
    products := []Product{
             {
               ID: 1,
               Name: "Laptop",
             },
             {
               ID: 2,
               Name: "Mouse",
             },
    }
    data, err = json.MarshalIndent(
                products,
                "",
                " ",
    )
    fmt.Println(string(data))
    fmt.Println("Guardar json en archivo")
    err = os.WriteFile(
          "products.json",
           data,
           0644,
    )
    content, err := os.ReadFile(
                  "products.json",
    )
    if err != nil {
      panic(err)
    }
    fmt.Println(string(content))
    fmt.Println("Deserializar Json")
    product = Product{
             ID: 32,
             Name: "Maria",
             Price: 984,
             Stock: 12,
    }
    data, err = json.MarshalIndent(
               product,
               "",
               " ",
    )
    if err != nil {
         panic(err)
    }
    err = os.WriteFile(
         "product.json",
          data,
          0644,
    )
    if err != nil {
          panic(err)
    }
    content, err = os.ReadFile(
         "product.json",
    )
    if err != nil {
          panic(err)
    }
    var loaded Product
    err = json.Unmarshal(
         content,
         &loaded,
    )
    if err != nil {
      panic(err)
    }
    fmt.Println(loaded.Name)
}
