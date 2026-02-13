package mcpx

import (
	"log/slog"
	"net/http"

	mlog "github.com/hilmoo/go-backend-example/internal/transport/middleware/log"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewStreamableHTTPHandler(s *mcp.Server) *mcp.StreamableHTTPHandler {
	return mcp.NewStreamableHTTPHandler(func(r *http.Request) *mcp.Server {
		mlog.InitMcp(r)
		return s
	}, &mcp.StreamableHTTPOptions{JSONResponse: true})
}

func NewServer(impl *mcp.Implementation, log *slog.Logger) *mcp.Server {
	s := mcp.NewServer(impl, nil)
	s.AddReceivingMiddleware(mlog.Mcp(log))
	return s
}
