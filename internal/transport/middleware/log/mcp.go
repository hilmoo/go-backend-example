package mlog

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Mcp(logger *slog.Logger) mcp.Middleware {
	return func(next mcp.MethodHandler) mcp.MethodHandler {
		return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
			start := time.Now()
			sessionID := req.GetSession().ID()
			var errCode int64
			var hasError bool

			defer func() {
				end := time.Now()
				latency := end.Sub(start)

				attrs := []slog.Attr{
					slog.String("method", method),
					slog.String("session_id", sessionID),
					slog.Duration("latency", latency),
				}

				// InitMcp
				if v := ctx.Value(reqMcpAttributesCtxKey); v != nil {
					v.(*sync.Map).Range(func(key, value any) bool {
						attrs = append(attrs, slog.Attr{Key: key.(string), Value: value.(slog.Value)})
						return true
					})
				}

				if v := ctx.Value(customAttributesCtxKey); v != nil {
					v.(*sync.Map).Range(func(key, value any) bool {
						k := key.(string)
						val := value.(slog.Value)

						// Check error for log level determination
						if k == "error" {
							hasError = true
							if val.Kind() == slog.KindGroup {
								for _, attr := range val.Group() {
									if attr.Key == "code" {
										errCode = attr.Value.Int64()
									}
								}
							}
						}
						attrs = append(attrs, slog.Attr{Key: k, Value: val})
						return true
					})
				}

				level := DefaultLevel
				if hasError {
					if errCode >= 500 {
						level = ServerErrorLevel
					} else if errCode >= 400 {
						level = ClientErrorLevel
					}
				}

				logger.LogAttrs(ctx, level, "mcp_request", attrs...)
			}()

			return next(ctx, method, req)
		}
	}
}

func InitMcp(r *http.Request) {
	if va := r.Context().Value(reqMcpAttributesCtxKey); va == nil {
		r = r.WithContext(context.WithValue(r.Context(), reqMcpAttributesCtxKey, &sync.Map{}))
	}

	ctx := r.Context()

	if v := ctx.Value(reqMcpAttributesCtxKey); v != nil {
		if m, ok := v.(*sync.Map); ok {
			m.Store("request_id", slog.StringValue(r.Header.Get(echo.HeaderXRequestID)))
			m.Store("remote_ip", slog.StringValue(r.RemoteAddr))
		}
	}
}
