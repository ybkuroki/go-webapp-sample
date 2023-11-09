package session

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/model"
	"gopkg.in/boj/redistore.v1"
)

const (
	// sessionStr represents a string of session key.
	sessionStr = "GSESSION"
	// Account is the key of account data in the session.
	Account = "Account"
)

type session struct {
	store sessions.Store
}

// Session represents a interface for accessing the session on the application.
type Session interface {
	GetStore() sessions.Store

	Get(c echo.Context) *sessions.Session
	Save(c echo.Context) error
	Delete(c echo.Context) error
	SetValue(c echo.Context, key string, value interface{}) error
	GetValue(c echo.Context, key string) string
	SetAccount(c echo.Context, account *model.Account) error
	GetAccount(c echo.Context) *model.Account
}

// NewSession is constructor.
func NewSession(logger logger.Logger, conf *config.Config) Session {
	if !conf.Redis.Enabled {
		logger.GetZapLogger().Infof("use CookieStore for session")
		return &session{sessions.NewCookieStore([]byte("secret"))}
	}

	logger.GetZapLogger().Infof("use redis for session")
	logger.GetZapLogger().Infof("Try redis connection")
	address := fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port)
	store, err := redistore.NewRediStore(conf.Redis.ConnectionPoolSize, "tcp", address, "", []byte("secret"))
	if err != nil {
		logger.GetZapLogger().Panicf("Failure redis connection, %s", err.Error())
	}
	logger.GetZapLogger().Infof(fmt.Sprintf("Success redis connection, %s", address))
	return &session{store: store}
}

func (s *session) GetStore() sessions.Store {
	return s.store
}

// Get returns a session for the current request.
func (s *session) Get(c echo.Context) *sessions.Session {
	sess, _ := s.store.Get(c.Request(), sessionStr)
	return sess
}

// Save saves the current session.
func (s *session) Save(c echo.Context) error {
	sess := s.Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
	return s.saveSession(c, sess)
}

// Delete the current session.
func (s *session) Delete(c echo.Context) error {
	sess := s.Get(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	}
	return s.saveSession(c, sess)
}

func (s *session) saveSession(c echo.Context, sess *sessions.Session) error {
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return fmt.Errorf("error occurred while save session")
	}
	return nil
}

// SetValue sets a key and a value.
func (s *session) SetValue(c echo.Context, key string, value interface{}) error {
	sess := s.Get(c)
	bytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json marshal error while set value in session")
	}
	sess.Values[key] = string(bytes)
	return nil
}

// GetValue returns value of session.
func (s *session) GetValue(c echo.Context, key string) string {
	sess := s.Get(c)
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

func (s *session) SetAccount(c echo.Context, account *model.Account) error {
	return s.SetValue(c, Account, account)
}

func (s *session) GetAccount(c echo.Context) *model.Account {
	if v := s.GetValue(c, Account); v != "" {
		a := &model.Account{}
		_ = json.Unmarshal([]byte(v), a)
		return a
	}
	return nil
}
