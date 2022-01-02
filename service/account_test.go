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
	assert.True(t, result)
}

func TestAuthenticateByUsernameAndPassword_EntityNotFound(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewAccountService(container)
	result, account := service.AuthenticateByUsernameAndPassword("abcde", "abcde")

	assert.Nil(t, account)
	assert.False(t, result)
}

func TestAuthenticateByUsernameAndPassword_AuthenticationFailure(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewAccountService(container)
	result, account := service.AuthenticateByUsernameAndPassword("test", "abcde")

	assert.Nil(t, account)
	assert.False(t, result)
}
