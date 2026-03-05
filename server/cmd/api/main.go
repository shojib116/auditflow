package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/shojib116/auditflow-api/config"
	iamService "github.com/shojib116/auditflow-api/internal/application/iam"
	"github.com/shojib116/auditflow-api/internal/database"
	iamRepository "github.com/shojib116/auditflow-api/internal/infra/iam"
	iamHandler "github.com/shojib116/auditflow-api/internal/interfaces/http/iam"
	"github.com/shojib116/auditflow-api/internal/interfaces/http/middlewares"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status: Healthy"))
}

func main() {
	cfg := config.GetConfig()

	dbURL := database.GetConnectionString(cfg.DB)
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err.Error())
	}

	iamRepo := iamRepository.NewUserRepository(dbConn)
	iamSvc := iamService.NewUserService(iamRepo, cfg)
	iamHndlr := iamHandler.NewHandler(&iamSvc)

	mngr := middlewares.NewManager(middlewares.Logger, middlewares.CORS(cfg.FrontendDomain))
	mux := http.NewServeMux()

	mux.Handle("GET /healthz", http.HandlerFunc(handlerHealth))

	iamHndlr.RegisterRoutes(mux, mngr)

	serverAddr := fmt.Sprintf(":%d", cfg.HttpPort)
	fmt.Printf("Server running on http://localhost%v\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, mngr.Wrap(mux)); err != nil {
		log.Fatal("Error occured:", err)
	}

}
