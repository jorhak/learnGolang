package repository

import (
      "encoding/json"
      "os"
      "erp-customer-storage/internal/domain"
)

type JSONRepository struct{
    FileName string
}

func NewJSONRepository(filename string) domain.CustomerRepository {
    return &JSONRepository{
          FileName: filename,
    }
}

func (m *JSONRepository) Store(customers []domain.Customer) error {
    data, err := json.MarshalIndent(
                     customers,
                     "",
                     " ",
                 )
    if err != nil {
      return err
    }
    return  os.WriteFile(
                 m.FileName,
                 data,
                 0644,
            )
}
