package todo

import (
	"context"
	"fmt"
	"log/slog"

	errort "github.com/hilmoo/go-backend-example/internal/transport/error"
	"github.com/hilmoo/go-backend-example/internal/transport/mcpx"
	"github.com/hilmoo/go-backend-example/internal/transport/validation"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type mcpHandler struct {
	log      *slog.Logger
	validate *validation.Vld
	db       DbRepository
}

func NewMCPHandler(log *slog.Logger, vld *validation.Vld, db DbRepository) *mcpHandler {
	return &mcpHandler{
		log:      log,
		validate: vld,
		db:       db,
	}
}

func (h *mcpHandler) Register(s *mcp.Server) {
	mcp.AddTool(s, &mcp.Tool{
		Name:        "todo_get",
		Description: "Retrieve a todo item by its ID.",
	}, h.getTodo)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "todo_create",
		Description: "Create a new todo item with the provided details.",
	}, h.createTodo)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "todo_update",
		Description: "Update an existing todo item with new details.",
	}, h.updateTodo)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "todo_delete",
		Description: "Delete a todo item by its ID.",
	}, h.deleteTodo)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "todo_list",
		Description: "List all todo items with optional filtering.",
	}, h.listTodos)
}

func (h *mcpHandler) getTodo(ctx context.Context, _ *mcp.CallToolRequest, args GetTodoRequest) (*mcp.CallToolResult, any, error) {
	err := validation.ValidatePayload(h.validate, &args)
	if err != nil {
		return errort.McpError(ctx, err)
	}

	resp, err := getTodo(ctx, &args, h.db)
	if err != nil {
		return errort.McpError(ctx, err)
	}

	title := fmt.Sprintf("Details for Todo ID: %s", args.ID)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: mcpx.ToMcpText(title, resp),
			},
		},
	}, resp, nil
}

func (h *mcpHandler) createTodo(ctx context.Context, _ *mcp.CallToolRequest, args CreateTodoRequest) (*mcp.CallToolResult, any, error) {
	err := validation.ValidatePayload(h.validate, &args)
	if err != nil {
		return errort.McpError(ctx, err)
	}

	resp, err := createTodo(ctx, &args, h.db)
	if err != nil {
		return errort.McpError(ctx, err)
	}

	title := fmt.Sprintf("Created Todo ID: %s", resp.ID)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: mcpx.ToMcpText(title, resp),
			},
		},
	}, resp, nil
}

func (h *mcpHandler) updateTodo(ctx context.Context, _ *mcp.CallToolRequest, args UpdateTodoRequest) (*mcp.CallToolResult, any, error) {
	err := validation.ValidatePayload(h.validate, &args)
	if err != nil {
		return errort.McpError(ctx, err)
	}

	resp, err := updateTodo(ctx, &args, h.db)
	if err != nil {
		return errort.McpError(ctx, err)
	}

	title := fmt.Sprintf("Updated Todo ID: %s", resp.ID)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: mcpx.ToMcpText(title, resp),
			},
		},
	}, resp, nil
}

func (h *mcpHandler) deleteTodo(ctx context.Context, _ *mcp.CallToolRequest, args DeleteTodoRequest) (*mcp.CallToolResult, any, error) {
	err := validation.ValidatePayload(h.validate, &args)
	if err != nil {
		return errort.McpError(ctx, err)
	}

	resp, err := deleteTodo(ctx, &args, h.db)
	if err != nil {
		return errort.McpError(ctx, err)
	}

	title := fmt.Sprintf("Deleted Todo ID: %s", args.ID)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: mcpx.ToMcpText(title, resp),
			},
		},
	}, resp, nil
}

func (h *mcpHandler) listTodos(ctx context.Context, _ *mcp.CallToolRequest, args ListTodosRequest) (*mcp.CallToolResult, any, error) {
	err := validation.ValidatePayload(h.validate, &args)
	if err != nil {
		return errort.McpError(ctx, err)
	}

	resp, err := listTodos(ctx, &args, h.db)
	if err != nil {
		return errort.McpError(ctx, err)
	}

	title := "List of Todos"
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: mcpx.ToMcpText(title, resp),
			},
		},
	}, resp, nil
}
