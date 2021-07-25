package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// FormatController is a controller for managing format data.
type FormatController struct {
	context mycontext.Context
	service *service.FormatService
}

// NewFormatController is constructor.
func NewFormatController(context mycontext.Context) *FormatController {
	return &FormatController{context: context, service: service.NewFormatService(context)}
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
func (controller *FormatController) GetFormatList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllFormats())
}
