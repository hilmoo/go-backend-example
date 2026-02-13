package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hilmoo/go-backend-example/internal/transport/validation"
	"github.com/labstack/echo/v5"
)

var Version = "dev"

func Main(cfg Config) error {
	logger := initLogger(cfg.LogLevel)
	vld := validation.InitValidation()

	db, err := initDb(cfg)
	if err != nil {
		logger.Error("failed to initialize database", slog.String("error", err.Error()))
		return fmt.Errorf("initialize database: %w", err)
	}
	defer db.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app := initHandler(logger, vld, db, cfg)
	sc := echo.StartConfig{
		Address:         fmt.Sprintf("%s:%d", cfg.ListenAddr, cfg.ListenPort),
		HideBanner:      true,
		GracefulTimeout: 10 * time.Second,
	}
	if err := sc.Start(ctx, app); err != nil {
		logger.Error("failed to start server", slog.String("error", err.Error()))
		return fmt.Errorf("start server: %w", err)
	}

	return nil
}
