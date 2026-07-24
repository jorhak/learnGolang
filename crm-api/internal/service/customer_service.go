package service

import(
   "log"
   "encoding/json"
   "crm-api/internal/domain"
)

// CustomerService contiene logica del negocio relacionada con los clientes
type CustomerService struct {}

// NewCustomerService crea una nueva instanica del servicio
func NewCustomerService() *CustomerService {
    return &CustomerService{}
}

// GetCustomers devuelve el listado de clientes
//
// Por ahora los datos son simulados
// En los proximos dias se obtendran desde repository
func (s *CustomerService) GetCustomers() []domain.Customer {
    customers := []domain.Customer{
                {
                   ID:    1,
                   Name:  "Jon Snow",
                   Email: "jon@snow.com",
                },
                {
                   ID:    2,
                   Name:  "Arya Stark",
                   Email: "arya@stark.com",
                },
    }

    return customers
}
// Funcion que busca el objeto con el ID
func (s *CustomerService) GetCustomerByID(id int) (domain.Customer, bool) {
    customers := s.GetCustomers()
    result := buscarPorID(id, customers)
    if result != nil {
      data, err := json.MarshalIndent(result, "","")
      if err != nil {
        log.Println("No se pudo serializar el objeto", err)
      }
      log.Printf("Objeto encontrado:\n%s", string(data))
      return *result, true
    } else {
      log.Println("No se encontro ningun objeto con ese ID:",id)
      return domain.Customer{}, false
    }
}
// Funcion que busca y devuelve el objeto encontrado
func buscarPorID(id int, customers []domain.Customer) *domain.Customer {
    for i := range customers {
        if customers[i].ID == id {
          return &customers[i]
        }
    }
    return nil
}
