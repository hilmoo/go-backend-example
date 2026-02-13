package cmd

import (
	"context"
	"fmt"

	"github.com/hilmoo/go-backend-example/internal/app"
	"github.com/urfave/cli/v3"
)

// versionCommand returns a CLI command that prints the application version.
func versionCommand() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Print the application version",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Printf("%s\n", app.Version)
			return nil
		},
	}
}
