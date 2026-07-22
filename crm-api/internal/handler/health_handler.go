package handler

import(
   "net/http"
)

type HealthHandler struct {}

func NewHealthHandler() *HealthHandler {
    return &HealthHandler{}
}

func (h *HealthHandler) Health(
    w http.ResponseWriter, r *http.Request,
){
    w.Header().Set(
        "Content-Type",
        "text/planin",
    )
    w.WriteHeader(http.StatusOK)
    _, _ = w.Write(
        []byte("CRM API Running"),
    )
}
