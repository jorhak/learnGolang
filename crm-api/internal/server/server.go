package server

import(
   "log"
   "net/http"
)

type Server struct {
    mux *http.ServeMux
}

func New() *Server {
	  mux := http.NewServeMux()
    s := &Server{
        mux: mux,
    }
		RegisterRoutes(
			mux,
		)
    return s
}

func (s *Server) Start() error {
    log.Println("Servidor escuchando en :8080")
    return http.ListenAndServe(
           ":8080",
           s.mux,
    )
}
