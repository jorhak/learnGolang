package server

import(
   "crm-api/internal/handler"
)

func (s *Server) routes() {
    healthHandler := handler.NewHealthHandler()
    versionHandler := handler.NewVersionHandler()

    s.mux.HandleFunc(
       "/health",
       healthHandler.Health,
    )

    s.mux.HandleFunc(
       "/version",
       versionHandler.Version,
    )
}
