package main

import "fmt"
import "encoding/json"

type Product struct {
    ID    int
    Name  string
    Price float64
    Stock int
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
}
