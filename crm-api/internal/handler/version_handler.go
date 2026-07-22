package handler

import(
   "net/http"
)

type VersionHandler struct{}

func NewVersionHandler() *VersionHandler {
    return &VersionHandler{}
}

func (h *VersionHandler) Version(
    w http.ResponseWriter,
    r *http.Request,
) {
    w.Header().Set(
      "Content-Type",
      "text/plain",
    )
    w.WriteHeader(http.StatusOK)
    _, err := w.Write([]byte("CRM API v1.0.0"))
    if err != nil {
        http.Error(
             w,
             "Error writing response",
             http.StatusInternalServerError,
        )
        return
    }
}
