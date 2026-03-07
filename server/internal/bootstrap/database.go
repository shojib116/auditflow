package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/shojib116/auditflow-api/config"
	"github.com/shojib116/auditflow-api/internal/database"
)

func setupDatabase(cfg *config.DBConfig) (*sql.DB, error) {
	dbURL := database.GetConnectionString(cfg)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return db, nil
}
