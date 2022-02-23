package service

import (
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
	"golang.org/x/crypto/bcrypt"
)

// AccountService is a service for managing user account.
type AccountService interface {
	AuthenticateByUsernameAndPassword(username string, password string) (bool, *model.Account)
}

type accountService struct {
	container container.Container
}

// NewAccountService is constructor.
func NewAccountService(container container.Container) AccountService {
	return &accountService{container: container}
}

// AuthenticateByUsernameAndPassword authenticates by using username and plain text password.
func (a *accountService) AuthenticateByUsernameAndPassword(username string, password string) (bool, *model.Account) {
	rep := a.container.GetRepository()
	logger := a.container.GetLogger()
	account := model.Account{}
	result, err := account.FindByName(rep, username)
	if err != nil {
		logger.GetZapLogger().Errorf(err.Error())
		return false, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
		return false, nil
	}

	return true, result
}
