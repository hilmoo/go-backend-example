package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	// Base
	ListenPort int    `env:"LISTEN_PORT" envDefault:"8080"`
	ListenAddr string `env:"LISTEN_ADDR" envDefault:"0.0.0.0"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"INFO"`

	EnableHealth bool `env:"ENABLE_HEALTH" envDefault:"true"`

	// Data
	DataFolder string `env:"DATA_FOLDER" envDefault:"data"`
	SqliteDSN string `env:"SQLITE_DSN"`
}

func LoadConfig() (Config, error) {
	var c Config
	if err := env.Parse(&c); err != nil {
		return c, fmt.Errorf("env config: %w", err)
	}

	if err := ensureDirExists(c.DataFolder); err != nil {
		return c, fmt.Errorf("ensure data dir exists: %w", err)
	}

	if c.SqliteDSN == "" {
		dbPath := filepath.Join(c.DataFolder, "app.db")
		c.SqliteDSN = fmt.Sprintf("file:%s", dbPath)
	}

	return c, nil
}

func ensureDirExists(dir string) error {
	if dir == "" || dir == "." {
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}
