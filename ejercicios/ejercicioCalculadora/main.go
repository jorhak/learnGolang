package main

import "fmt"

func sum(a, b float64) float64{
    return a + b
}

func subtract(a, b float64) float64{
    return a - b
}

func main(){
    total := sum(100,50)
    discount :=subtract(total, 10)

    fmt.Println("Total:", total)
    fmt.Println("Descuento:", discount)
}
