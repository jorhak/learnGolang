package domain

type CustomerRepository interface {
    Store([]Customer) error
    Load() ([]Customer, error)
}
