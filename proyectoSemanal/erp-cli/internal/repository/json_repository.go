package repository

import (
   "encoding/json"
   "os"
   "proyectoSemanal/erp-cli/internal/domain"
)

type JSONRepository struct {
    FileName string
}

func NewJSONRepository(
    filename string,
) *JSONRepository {
  return &JSONRepository{
            FileName: filename,
  }
}

func (j *JSONRepository) Save(products []domain.Product) error {
    data, err := json.MarshalIndent(
                       products,
                       "",
                       " ",
    )

    if err != nil {
      return err
    }

    return os.WriteFile(
               j.FileName,
               data,
               0644,
    )
}

func (j *JSONRepository) Load() ([]domain.Product, error) {
    data, err := os.ReadFile(j.FileName)

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

