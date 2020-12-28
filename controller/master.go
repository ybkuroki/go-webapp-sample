package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// MasterController is a controller for managing master data such as format and category.
type MasterController struct {
	context mycontext.Context
	service *service.MasterService
}

// NewMasterController is constructor.
func NewMasterController(context mycontext.Context) *MasterController {
	return &MasterController{context: context, service: service.NewMasterService(context)}
}

// GetCategoryList returns the list of all categories.
func (controller *MasterController) GetCategoryList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllCategories())
}

// GetFormatList returns the list of all formats.
func (controller *MasterController) GetFormatList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllFormats())
}
