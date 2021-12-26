package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestAuthenticateByUsernameAndPassword_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewAccountService(container)
	result, account := service.AuthenticateByUsernameAndPassword("test", "test")

	a := model.Account{}
	data, _ := a.FindByName(container.GetRepository(), "test")

	assert.Equal(t, data, account)
	assert.Equal(t, true, result)
}

func TestAuthenticateByUsernameAndPassword_EntityNotFound(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewAccountService(container)
	result, account := service.AuthenticateByUsernameAndPassword("abcde", "abcde")

	assert.Equal(t, (*model.Account)(nil), account)
	assert.Equal(t, false, result)
}

func TestAuthenticateByUsernameAndPassword_AuthenticationFailure(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewAccountService(container)
	result, account := service.AuthenticateByUsernameAndPassword("test", "abcde")

	assert.Equal(t, (*model.Account)(nil), account)
	assert.Equal(t, false, result)
}
