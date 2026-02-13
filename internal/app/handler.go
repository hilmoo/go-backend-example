package app

import (
	"database/sql"
	"log/slog"

	"github.com/hilmoo/go-backend-example/doc"
	"github.com/hilmoo/go-backend-example/internal/feature/healthz"
	"github.com/hilmoo/go-backend-example/internal/feature/todo"
	"github.com/hilmoo/go-backend-example/internal/store/db"
	"github.com/hilmoo/go-backend-example/internal/transport/mcpx"
	logx "github.com/hilmoo/go-backend-example/internal/transport/middleware/log"
	"github.com/hilmoo/go-backend-example/internal/transport/validation"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func initHandler(logger *slog.Logger, vld *validation.Vld, dbConn *sql.DB, cfg Config) *echo.Echo {
	e := echo.New()

	todoRepo := db.NewTodoRepository(dbConn)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID(), middleware.CORS("*"), logx.New(logger).EchoMiddleware(), middleware.Recover())

	apiGroup := e.Group("/api")
	{
		registerSwagger(apiGroup)
		todo.NewHTTPHandler(vld, todoRepo).Register(apiGroup)
	}

	setupMCP(e, logger, vld, todoRepo)

	if cfg.EnableHealth {
		healthz.NewHTTPHandler().Register(e.Group(""))
	}

	return e
}

func registerSwagger(g *echo.Group) {
	g.GET("/swagger/doc.json", func(c *echo.Context) error {
		return c.Blob(200, "application/json", doc.SwaggerJSON)
	})
	// TODO: https://github.com/swaggo/echo-swagger/pull/140
}

func setupMCP(e *echo.Echo, logger *slog.Logger, vld *validation.Vld, repo todo.DbRepository) {
	mcpServer := mcpx.NewServer(&mcp.Implementation{
		Name:    "Go Backend Example",
		Version: "1.0.0",
	}, logger)

	todo.NewMCPHandler(logger, vld, repo).Register(mcpServer)

	mcpHandler := mcpx.NewStreamableHTTPHandler(mcpServer)
	e.Group("/mcp").Any("", echo.WrapHandler(mcpHandler))
}
