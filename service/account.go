package service

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticateByUsernameAndPassword authenticates by using username and plain text password.
func AuthenticateByUsernameAndPassword(context mycontext.Context, username string, password string) (bool, *model.Account) {
	rep := context.GetRepository()
	logger := context.GetLogger()
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
