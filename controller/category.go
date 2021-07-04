package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// CategoryController is a controller for managing master data such as format and category.
type CategoryController struct {
	context mycontext.Context
	service *service.CategoryService
}

// NewCategoryController is constructor.
func NewCategoryController(context mycontext.Context) *CategoryController {
	return &CategoryController{context: context, service: service.NewCategoryService(context)}
}

// GetCategoryList returns the list of all categories.
func (controller *CategoryController) GetCategoryList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllCategories())
}
