package domain

type ProductRepository interface {
    Save(products []Product) error
    Load() ([]Product, error)
}
