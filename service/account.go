package service

import (
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticateByUsernameAndPassword authenticates by using username and plain text password.
func AuthenticateByUsernameAndPassword(username string, password string) (bool, *model.Account) {
	rep := repository.GetRepository()
	account := model.Account{}
	result, err := account.FindByName(rep, username)
	if err != nil {
		logger.GetEchoLogger().Error(err)
		return false, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		logger.GetEchoLogger().Error(err)
		return false, nil
	}

	return true, result
}
