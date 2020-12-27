package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
)

// HealthController is
type HealthController struct {
	context mycontext.Context
}

// NewHealthController is
func NewHealthController(context mycontext.Context) *HealthController {
	return &HealthController{context: context}
}

// GetHealthCheck returns whether this application is alive or not.
func (controller *HealthController) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}
