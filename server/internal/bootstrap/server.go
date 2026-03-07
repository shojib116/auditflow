package bootstrap

import (
	"context"
	"fmt"
	iamHandler "github.com/shojib116/auditflow-api/internal/interfaces/http/iam"
	"github.com/shojib116/auditflow-api/internal/interfaces/http/middlewares"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	port int
	mngr *middlewares.Manager
	http *http.Server

	// handlers — add more as your app grows
	iamHandler *iamHandler.Handler
}

func newServer(port int, mngr *middlewares.Manager, iamHndlr *iamHandler.Handler) *Server {
	s := &Server{port: port, mngr: mngr, iamHandler: iamHndlr}

	mux := http.NewServeMux()
	s.registerRoutes(mux)

	s.http = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mngr.Wrap(mux),
	}

	return s
}

func (s *Server) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		fmt.Printf("Server running on http://localhost:%d\n", s.port)
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		fmt.Println("shutting down...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return s.http.Shutdown(shutdownCtx)
	}
}
