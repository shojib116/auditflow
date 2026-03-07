package bootstrap

import "net/http"

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.Handle("GET /healthz", http.HandlerFunc(handleHealth))
	s.iamHandler.RegisterRoutes(mux, s.mngr)
	// future: s.auditHandler.RegisterRoutes(mux, s.mngr)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status: Healthy"))
}
