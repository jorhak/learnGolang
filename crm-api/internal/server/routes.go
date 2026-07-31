package server

import(
   "crm-api/internal/handler"
   "crm-api/internal/service"
	 "crm-api/internal/repository"
	 "net/http"
)

func RegisterRoutes(
	mux *http.ServeMux,
) {

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

    mux.HandleFunc(
       "/health",
       healthHandler.Health,
    )

    mux.HandleFunc(
       "/version",
       versionHandler.Version,
    )

    mux.HandleFunc(
       "/customers",
			 func (
				 w http.ResponseWriter,
				 r *http.Request,
			 ) {
				 switch r.Method {

				 case http.MethodGet:
					 customerHandler.GetCustomers(
						 w,
						 r,
					 )

				 case http.MethodPost:
					 customerHandler.CreateCustomer(
						 w,
						 r,
					 )

				 default:
					 http.NotFound(
						 w,
						 r,
					 )

				}
			 },
    )

		mux.HandleFunc(
			"/delete",
			customerHandler.DeleteCustomer,
		)

    mux.HandleFunc(
       "/customerResponse",
       customerResponse.GetCustomerResponse,
    )

    mux.HandleFunc(
       "/customer",
       customer.Get,
    )

    mux.HandleFunc(
       "/cliente",
       customerHandler.GetCustomerByID,
    )
}
