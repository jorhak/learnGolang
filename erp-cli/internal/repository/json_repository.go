package repository

import (
  "encoding/json"
  "os"
  "erp-cli/internal/domain"
)

type JSONRepository struct {
    FileName string
}

func (r JSONRepository) Save(products []domain.Product) error {
    data, err := json.MarshalIndent(
              products,
              "",
              " ",
    )
    if err != nil {
           return err
    }
    return os.WriteFile(
          r.FileName,
          data,
          0644,
    )
}

func (r JSONRepository) Load() ([]domain.Product, error) {
    data, err := os.ReadFile(
             r.FileName,
    )
    if err != nil {
          return nil, err
    }
    var products []domain.Product
    err = json.Unmarshal(
       data,
       &products,
    )
    if err != nil {
          return nil, err
    }
    return products, nil
}


