package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Config struct {
	DSN string
}

func (cfg *Config) NewConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}