package server

import(
   "crm-api/internal/handler"
)

func (s *Server) routes() {
    healthHandler := handler.NewHealthHandler()
    versionHandler := handler.NewVersionHandler()
    customerHandler := handler.NewCustomerHandler()

    s.mux.HandleFunc(
       "/health",
       healthHandler.Health,
    )

    s.mux.HandleFunc(
       "/version",
       versionHandler.Version,
    )

    s.mux.HandleFunc(
       "/customers",
       customerHandler.GetCustomers,
    )
}
