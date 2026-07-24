package handler

import(
   "log"
   "net/http"
   "encoding/json"
)

type CustomerResponse struct {
    Message string `json:"message"`
}

func NewCustomerResponse(
    message string,
) *CustomerResponse {
    return &CustomerResponse{
       Message: message,
    }
}

func (c *CustomerResponse) GetCustomerResponse(
    w http.ResponseWriter,
    r *http.Request,
) {

  //response := c.Message

  data, err := json.MarshalIndent(
       c,
       "",
       " ",
  )

  if err != nil {
    http.Error(w, "Internal Error", http.StatusInternalServerError)
    return
  }

  w.Header().Set(
    "Content-Type",
    "application/json",
  )

  w.WriteHeader(http.StatusOK)
  _, err = w.Write(data)
  if err != nil {
    log.Fatal("No se obtuvo la data", err)
    return
  }
  log.Printf("Se agrego:\n%s",string(data))

}

