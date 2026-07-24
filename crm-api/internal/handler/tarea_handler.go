package handler

import(
   "log"
   "encoding/json"
   "net/http"
)

type Customer struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Email string  `json:"email"`
}

type TareaHandler struct {}

func NewTareaHandler() *TareaHandler {
    return &TareaHandler{}
}

func (t *TareaHandler) Get(
    w http.ResponseWriter,
    r *http.Request,
) {
    customers := []Customer{
              {
                  ID: 1,
                  Name: "Jon Snow",
                  Email: "jon@example.com",
              },
              {
                  ID:2,
                  Name: "Arya Stark",
                  Email: "arya@examle.com",
              },
    }
    data, err := json.MarshalIndent(
              customers,
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
    w.Header().Set(
       "Author",
       "KrekenD",
    )

    w.WriteHeader(http.StatusOK)
    _, err = w.Write(data)
    if err != nil {
       log.Fatal("Error en la serializacion de los datos", err)
       return
    }
    log.Printf(
        "Respuesta: \n%s",
        data,
    )
}

