package domain

type Order struct {
    ID int
    Product Product
    Customer Customer
    Quantity int
}

func (o Order) Total() float64 {
    return o.Product.Price * float64(o.Quantity)
}
