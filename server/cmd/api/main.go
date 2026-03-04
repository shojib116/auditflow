package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shojib116/auditflow-api/config"
	"github.com/shojib116/auditflow-api/internal/interfaces/http/middlewares"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status: Healthy"))
}

func main() {
	cfg := config.GetConfig()

	mngr := middlewares.NewManager(middlewares.Logger, middlewares.CORS(cfg.FrontendDomain))
	mux := http.NewServeMux()

	mux.Handle("GET /healthz", http.HandlerFunc(handlerHealth))

	serverAddr := fmt.Sprintf(":%d", cfg.HttpPort)
	fmt.Printf("Server running on http://localhost%v", serverAddr)
	if err := http.ListenAndServe(serverAddr, mngr.Wrap(mux)); err != nil {
		log.Fatal("Error occured:", err)
	}

}
