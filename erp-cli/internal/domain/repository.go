package domain

type ProductRepository interface {
    Save([]Product) error
    Load() ([]Product, error)
}
