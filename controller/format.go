package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// FormatController is a controller for managing master data such as format and category.
type FormatController struct {
	context mycontext.Context
	service *service.FormatService
}

// NewFormatController is constructor.
func NewFormatController(context mycontext.Context) *FormatController {
	return &FormatController{context: context, service: service.NewFormatService(context)}
}

// GetFormatList returns the list of all formats.
func (controller *FormatController) GetFormatList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllFormats())
}
