package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status: Healthy"))
}

func main() {
	const PORT = "8080"
	mux := http.NewServeMux()

	mux.Handle("GET /healthz", http.HandlerFunc(handlerHealth))

	fmt.Printf("Server running on http://localhost:%v", PORT)
	if err := http.ListenAndServe(":"+PORT, mux); err != nil {
		log.Fatal("Error occured:", err)
	}

}
