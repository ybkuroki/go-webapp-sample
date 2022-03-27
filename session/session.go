package session

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	echoSession "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/model"
)

const (
	// sessionStr represents a string of session key.
	sessionStr = "GSESSION"
	// Account is the key of account data in the session.
	Account = "Account"
)

type session struct {
	context echo.Context
}

type Session interface {
	SetContext(c echo.Context)
	Get() *sessions.Session
	Save() error
	Delete() error
	SetValue(key string, value interface{}) error
	GetValue(key string) string
	SetAccount(account *model.Account) error
	GetAccount() *model.Account
}

func NewSession() Session {
	return &session{context: nil}
}

func (s *session) SetContext(c echo.Context) {
	s.context = c
}

// Get returns a session for the current request.
func (s *session) Get() *sessions.Session {
	if s.context != nil {
		sess, _ := echoSession.Get(sessionStr, s.context)
		return sess
	}
	return nil
}

// Save saves the current session.
func (s *session) Save() error {
	sess := s.Get()
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
	return s.saveSession(sess)
}

// Delete the current session.
func (s *session) Delete() error {
	sess := s.Get()
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	return s.saveSession(sess)
}

func (s *session) saveSession(sess *sessions.Session) error {
	if err := sess.Save(s.context.Request(), s.context.Response()); err != nil {
		return s.context.NoContent(http.StatusInternalServerError)
	}
	return nil
}

// SetValue sets a key and a value.
func (s *session) SetValue(key string, value interface{}) error {
	sess := s.Get()
	bytes, err := json.Marshal(value)
	if err != nil {
		return s.context.NoContent(http.StatusInternalServerError)
	}
	sess.Values[key] = string(bytes)
	return nil
}

// GetValue returns value of session.
func (s *session) GetValue(key string) string {
	sess := s.Get()
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
func (s *session) SetAccount(account *model.Account) error {
	return s.SetValue(Account, account)
}

// GetAccount returns account object of session.
func (s *session) GetAccount() *model.Account {
	if v := s.GetValue(Account); v != "" {
		a := &model.Account{}
		_ = json.Unmarshal([]byte(v), a)
		return a
	}
	return nil
}
