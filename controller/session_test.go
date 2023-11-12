package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/test"
)

type sessionController struct {
	container container.Container
}

func TestSessionRace_Success(t *testing.T) {
	sessionKey := "Key"
	router, container := test.PrepareForControllerTest(true)
	session := sessionController{container: container}

	router.GET(config.API+"1", func(c echo.Context) error {
		_ = session.container.GetSession().SetValue(c, sessionKey, 1)
		_ = session.container.GetSession().Save(c)
		time.Sleep(3 * time.Second)
		return c.String(http.StatusOK, session.container.GetSession().GetValue(c, sessionKey))
	})
	router.GET(config.API+"2", func(c echo.Context) error {
		_ = session.container.GetSession().SetValue(c, sessionKey, 2)
		_ = session.container.GetSession().Save(c)
		return c.String(http.StatusOK, session.container.GetSession().GetValue(c, sessionKey))
	})

	req1 := httptest.NewRequest("GET", config.API+"1", nil)
	req2 := httptest.NewRequest("GET", config.API+"2", nil)
	rec1 := httptest.NewRecorder()
	rec2 := httptest.NewRecorder()

	go func() {
		router.ServeHTTP(rec1, req1)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		router.ServeHTTP(rec2, req2)
	}()

	time.Sleep(5 * time.Second)

	assert.Equal(t, "1", rec1.Body.String())
	assert.Equal(t, "2", rec2.Body.String())
}
