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
}
