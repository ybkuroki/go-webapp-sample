package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/service"
)

// AccountController is a controller for managing user account.
type AccountController interface {
	GetLoginStatus(c echo.Context) error
	GetLoginAccount(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type accountController struct {
	context      container.Container
	service      service.AccountService
	dummyAccount *model.Account
}

// NewAccountController is constructor.
func NewAccountController(container container.Container) AccountController {
	return &accountController{
		context:      container,
		service:      service.NewAccountService(container),
		dummyAccount: model.NewAccountWithPlainPassword("test", "test", 1),
	}
}

// GetLoginStatus returns the status of login.
// @Summary Get the login status.
// @Description Get the login status of current logged-in user.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {boolean} bool "The current user have already logged-in. Returns true."
// @Failure 401 {boolean} bool "The current user haven't logged-in yet. Returns false."
// @Router /auth/loginStatus [get]
func (controller *accountController) GetLoginStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}

// GetLoginAccount returns the account data of logged in user.
// @Summary Get the account data of logged-in user.
// @Description Get the account data of logged-in user.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Account "Success to fetch the account data. If the security function is disable, it returns the dummy data."
// @Failure 401 {boolean} bool "The current user haven't logged-in yet. Returns false."
// @Router /auth/loginAccount [get]
func (controller *accountController) GetLoginAccount(c echo.Context) error {
	if !controller.context.GetConfig().Extension.SecurityEnabled {
		return c.JSON(http.StatusOK, controller.dummyAccount)
	}
	return c.JSON(http.StatusOK, controller.context.GetSession().GetAccount(c))
}

// Login is the method to login using username and password by http post.
// @Summary Login using username and password.
// @Description Login using username and password.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param data body dto.LoginDto true "User name and Password for logged-in."
// @Success 200 {object} model.Account "Success to the authentication."
// @Failure 401 {boolean} bool "Failed to the authentication."
// @Router /auth/login [post]
func (controller *accountController) Login(c echo.Context) error {
	dto := dto.NewLoginDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}

	sess := controller.context.GetSession()
	if account := sess.GetAccount(c); account != nil {
		return c.JSON(http.StatusOK, account)
	}

	authenticate, a := controller.service.AuthenticateByUsernameAndPassword(dto.UserName, dto.Password)
	if authenticate {
		_ = sess.SetAccount(c, a)
		_ = sess.Save(c)
		return c.JSON(http.StatusOK, a)
	}
	return c.NoContent(http.StatusUnauthorized)
}

// Logout is the method to logout by http post.
// @Summary Logout.
// @Description Logout.
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200
// @Router /auth/logout [post]
func (controller *accountController) Logout(c echo.Context) error {
	sess := controller.context.GetSession()
	_ = sess.SetAccount(c, nil)
	_ = sess.Delete(c)
	return c.NoContent(http.StatusOK)
}
