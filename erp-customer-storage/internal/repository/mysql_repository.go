package repository

import (
      "encoding/json"
      "os"
      "erp-customer-storage/internal/domain"
)

type MYSQLRepository struct {
    FileName string
}

func NewMYSQLRepository(filename string) domain.CustomerRepository {
    return &MYSQLRepository{
          FileName: filename,
    }
}

func (m MYSQLRepository) Store(customers []domain.Customer) error {
    data, err := json.MarshalIndent(
                      customers,
                      "",
                      " ",
                 )
    if err != nil {
      return err
    }
    return os.WriteFile(
          m.FileName,
          data,
          0644,
    )
}
