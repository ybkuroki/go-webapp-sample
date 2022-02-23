package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// CategoryController is a controller for managing category data.
type CategoryController interface {
	GetCategoryList(c echo.Context) error
}

type categoryController struct {
	container container.Container
	service   service.CategoryService
}

// NewCategoryController is constructor.
func NewCategoryController(container container.Container) CategoryController {
	return &categoryController{container: container, service: service.NewCategoryService(container)}
}

// GetCategoryList returns the list of all categories.
// @Summary Get a category list
// @Description Get a category list
// @Tags Categories
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Category "Success to fetch a category list."
// @Failure 401 {string} false "Failed to the authentication."
// @Router /categories [get]
func (controller *categoryController) GetCategoryList(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.FindAllCategories())
}
