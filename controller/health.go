package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/container"
)

// HealthController is a controller returns the current status of this application.
type HealthController interface {
	GetHealthCheck(c echo.Context) error
}

type healthController struct {
	container container.Container
}

// NewHealthController is constructor.
func NewHealthController(container container.Container) HealthController {
	return &healthController{container: container}
}

// GetHealthCheck returns whether this application is alive or not.
// @Summary Get the status of this application
// @Description Get the status of this application
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {string} message "healthy: This application is started."
// @Failure 404 {string} message "None: This application is stopped."
// @Router /health [get]
func (controller *healthController) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}
