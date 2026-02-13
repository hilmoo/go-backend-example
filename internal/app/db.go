package app

import (
	"database/sql"

	"github.com/hilmoo/go-backend-example/internal/store/db"
	_ "modernc.org/sqlite"
)

func initDb(cfg Config) (*sql.DB, error) {
	sqlite := db.Config{
		DSN: cfg.SqliteDSN,
	}

	dbConn, err := sqlite.NewConnection()
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
