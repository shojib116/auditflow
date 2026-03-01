package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shojib116/auditflow-api/internal/web/middlewares"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status: Healthy"))
}

func main() {
	const PORT = "8080"
	const ORIGIN = "http://localhost:5173"

	mngr := middlewares.NewManager(middlewares.Logger, middlewares.CORS(ORIGIN))
	mux := http.NewServeMux()

	mux.Handle("GET /healthz", http.HandlerFunc(handlerHealth))

	fmt.Printf("Server running on http://localhost:%v", PORT)
	if err := http.ListenAndServe(":"+PORT, mngr.Wrap(mux)); err != nil {
		log.Fatal("Error occured:", err)
	}

}
