package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hilmoo/go-backend-example/internal/app"
	"github.com/urfave/cli/v3"
)

func healthCheckCommand(cfg app.Config) *cli.Command {
	return &cli.Command{
		Name:  "healthcheck",
		Usage: "Run a health check against the application",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return runHealthCheck(cfg)
		},
	}
}

func runHealthCheck(cfg app.Config) error {
	if cfg.EnableHealth == false {
		return fmt.Errorf("health check endpoint is disabled")
	}

	host := cfg.ListenAddr
	if host == "0.0.0.0" {
		host = "localhost"
	}
	endpoint := fmt.Sprintf("http://%s:%d/api/healthz", host, cfg.ListenPort)

	resp, err := http.Get(endpoint)
	if err != nil {
		return fmt.Errorf("health check request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("health check failed with status code: %d", resp.StatusCode)
	}

	fmt.Println("Health check passed!")
	return nil
}
