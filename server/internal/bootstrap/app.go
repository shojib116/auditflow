package bootstrap

import (
	_ "github.com/lib/pq"
	"github.com/shojib116/auditflow-api/config"
	iamService "github.com/shojib116/auditflow-api/internal/application/iam"
	iamRepository "github.com/shojib116/auditflow-api/internal/infra/iam"
	iamHandler "github.com/shojib116/auditflow-api/internal/interfaces/http/iam"
	"github.com/shojib116/auditflow-api/internal/interfaces/http/middlewares"
)

type App struct {
	server *Server
}

func NewApp(cfg *config.Config) (*App, error) {
	// 1. infrastructure
	db, err := setupDatabase(cfg.DB)
	if err != nil {
		return nil, err
	}

	// 2. repositories
	iamRepo := iamRepository.NewUserRepository(db)

	// 3. services
	iamSvc := iamService.NewUserService(iamRepo, cfg)

	// 4. handlers
	iamHndlr := iamHandler.NewHandler(&iamSvc)

	// 5. http
	mngr := middlewares.NewManager(
		middlewares.Logger,
		middlewares.CORS(cfg.FrontendDomain),
	)
	server := newServer(cfg.HttpPort, mngr, iamHndlr)

	return &App{server: server}, nil
}

func (a *App) Run() error {
	return a.server.Start()
}
