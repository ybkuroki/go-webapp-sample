package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// FormatController is a controller for managing format data.
type FormatController interface {
	GetFormatList(c echo.Context) error
}

type formatController struct {
	container container.Container
	service   service.FormatService
}

// NewFormatController is constructor.
func NewFormatController(container container.Container) FormatController {
	return &formatController{container: container, service: service.NewFormatService(container)}
}

// GetFormatList returns the list of all formats.
// @Summary Get a format list
// @Description Get a format list
// @Tags Formats
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Format "Success to fetch a format list."
// @Failure 401 {string} false "Failed to the authentication."
// @Router /formats [get]
func (controller *formatController) GetFormatList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllFormats())
}
