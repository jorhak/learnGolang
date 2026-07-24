package handler

import(
   "fmt"
   "net/http"
)

type CustomerHandler struct {}

func NewCustomerHandler() *CustomerHandler{
    return &CustomerHandler{}
}

func (h *CustomerHandler) GetCustomers(
    w http.ResponseWriter,
    r *http.Request,
) {
    fmt.Println("Metodo:", r.Method)
    //Leer Url
    fmt.Println(r.URL.Path,)
    //Leer Query Parameters
    page := r.URL.Query().Get("page")
    limit := r.URL.Query().Get("limit")
    fmt.Println(page)
    fmt.Println(limit)
    //Leer Headers
    token := r.Header.Get("Authorization",)
    fmt.Println(token)
    //Enviar Header
    w.Header().Set(
       "Content-Type",
       "application/json",
    )
    w.Header().Set(
      "Author",
      "SABOR-LATINO",
    )
    //Enviar Status
    w.WriteHeader(http.StatusCreated)
    //w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("Cambio Listado de Clientes"))
}
