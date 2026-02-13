package healthz

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type httpHandler struct {
}

func NewHTTPHandler() *httpHandler {
	return &httpHandler{}
}

func (h *httpHandler) Register(e *echo.Group) {
	healthGroup := e.Group("/healthz")
	healthGroup.GET("", h.healthz)
}

func (h *httpHandler) healthz(c *echo.Context) error {
	return c.NoContent(http.StatusOK)
}
