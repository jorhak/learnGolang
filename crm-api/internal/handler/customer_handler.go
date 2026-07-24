package handler

import(
   "fmt"
   "net/http"
   "encoding/json"
   "strconv"
   "crm-api/internal/service"
)

// CustomerHandler maneja las peticiones HTTP relacionadas con clientes
type CustomerHandler struct {
    service *service.CustomerService
}

// NewCustomerHandler crea una nueva instancia del handler
func NewCustomerHandler(
    service *service.CustomerService,
) *CustomerHandler{
    return &CustomerHandler{
             service: service,
    }
}

// GetCustomers obtiene el listado de clientes
func (h *CustomerHandler) GetCustomers(
    w http.ResponseWriter,
    r *http.Request,
) {
    customers := h.service.GetCustomers()
    fmt.Println("Customer con Service")

    w.Header().Set(
       "Content-Type",
       "application/json",
    )
    w.Header().Set(
      "Author",
      "Maria-Joaquina",
    )

    w.WriteHeader(http.StatusOK)
    err := json.NewEncoder(w).Encode(customers)
    if err != nil {
      http.Error(
         w,
         "Internal Server Error",
         http.StatusInternalServerError,
      )
      return
    }
}

func (h *CustomerHandler) GetCustomerByID(
    w http.ResponseWriter,
    r *http.Request,
) {

    w.Header().Set(
       "Content-Type",
       "application/json",
    )
    w.Header().Set(
      "Author",
      "Maria-Joaquina",
    )
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
      http.Error(w, "El ID proporcionado debe ser un numero entero valido", http.StatusBadRequest)
      return
    }

    customer, exits := h.service.GetCustomerByID(id)
    if exits == false {
      http.Error(
        w,
        "Cliente no encontrado",
        http.StatusNotFound,
      )
      return
    }
    fmt.Println("Obteniendo customer por ID")

    w.WriteHeader(http.StatusOK)
    err = json.NewEncoder(w).Encode(customer)
    if err != nil {
      http.Error(
         w,
         "Internal Server Error",
         http.StatusInternalServerError,
      )
      return
    }
}
