package session

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/model"
)

const (
	// sessionStr represents a string of session key.
	sessionStr = "GSESSION"
	// Account is the key of account data in the session.
	Account = "Account"
)

// Get returns a session for the current request.
func Get(c echo.Context) *sessions.Session {
	sess, _ := session.Get(sessionStr, c)
	return sess
}

// Save saves the current session.
func Save(c echo.Context) error {
	sess := Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
	return saveSession(c, sess)
}

// Delete the current session.
func Delete(c echo.Context) error {
	sess := Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	return saveSession(c, sess)
}

func saveSession(c echo.Context, sess *sessions.Session) error {
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return nil
}

// SetValue sets a key and a value.
func SetValue(c echo.Context, key string, value interface{}) error {
	sess := Get(c)
	bytes, err := json.Marshal(value)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	sess.Values[key] = string(bytes)
	return nil
}

// GetValue returns value of session.
func GetValue(c echo.Context, key string) string {
	sess := Get(c)
	if sess != nil {
		if v, ok := sess.Values[key]; ok {
			data, result := v.(string)
			if result && data != "null" {
				return data
			}
		}
	}
	return ""
}

// SetAccount sets account data in session.
func SetAccount(c echo.Context, account *model.Account) error {
	return SetValue(c, Account, account)
}

// GetAccount returns account object of session.
func GetAccount(c echo.Context) *model.Account {
	v := GetValue(c, Account)
	if v != "" {
		a := &model.Account{}
		_ = json.Unmarshal([]byte(v), a)
		return a
	}
	return nil
}
