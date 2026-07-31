package handler

import(
   "fmt"
   "net/http"
   "encoding/json"
   "strconv"
	 "strings"
   "crm-api/internal/service"
	 "crm-api/internal/domain"
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

// Obtiene customer por medio de un ID
func (h *CustomerHandler) GetCustomerByID(
    w http.ResponseWriter,
    r *http.Request,
) {

		path := strings.TrimPrefix(
			r.URL.Path,
			"/customers/",
		)
    
    id, err := strconv.Atoi(path)

    if err != nil {
      http.Error(
				w,
				"El ID proporcionado debe ser un numero entero valido",
				http.StatusBadRequest,
			)
      return
    }

    customer, err := h.service.GetCustomerByID(id)

    if err != nil {
      http.Error(
        w,
        err.Error(),
        http.StatusNotFound,
      )
      return
    }
    fmt.Println("Obteniendo customer por ID")

		response, _ := json.Marshal(customer)

    w.Header().Set(
       "Content-Type",
       "application/json",
    )
    w.Header().Set(
      "Author",
      "Maria-Joaquina",
    )

		w.Write(response)
}

// Crea un customer
func (h *CustomerHandler) CreateCustomer(
	w http.ResponseWriter,
	r *http.Request,
) {
	var customer domain.Customer
	
	err := json.NewDecoder(
		r.Body,
	).Decode(&customer)

	if err != nil {
		http.Error(
			w,
			"Invalid Request",
			http.StatusBadRequest,
		)
		return
	}

	err = h.service.CreateCustomer(customer)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	w.WriteHeader(
		http.StatusCreated,
	)
}

// Elimina un customer por medio de ID
func (h *CustomerHandler) DeleteCustomer(
	w http.ResponseWriter,
	r *http.Request,
) {

	idStr := r.URL.Query().Get("id")
  id, err := strconv.Atoi(idStr)
  
	if err != nil {
    http.Error(w, "El ID proporcionado debe ser un numero entero valido", http.StatusBadRequest)
    return
  }

	err = h.service.DeleteCustomer(id)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}
	w.WriteHeader(
		http.StatusOK,
	)
}
