package server

import(
   "crm-api/internal/handler"
   "crm-api/internal/service"
	 "crm-api/internal/repository"
)

func (s *Server) routes() {
    healthHandler := handler.NewHealthHandler()
    versionHandler := handler.NewVersionHandler()
    customerResponse := handler.NewCustomerResponse("Saludos Terricolas")
    customer := handler.NewTareaHandler()

		memoryRepository := repository.NewMemoryCustomerRepository()
    customerService := service.NewCustomerService(
			memoryRepository,
		)
    customerHandler := handler.NewCustomerHandler(
                       customerService,
    )

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

    s.mux.HandleFunc(
       "/customerResponse",
       customerResponse.GetCustomerResponse,
    )

    s.mux.HandleFunc(
       "/customer",
       customer.Get,
    )

    s.mux.HandleFunc(
       "/cliente",
       customerHandler.GetCustomerByID,
    )
}
