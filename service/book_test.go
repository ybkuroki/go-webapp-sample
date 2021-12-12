package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestFindByID_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.FindByID("1")

	assert.Equal(t, uint(1), result.ID)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, result)
}

func TestFindByID_IdNotNumeric(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.FindByID("ABCD")

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, errors.New("failed to fetch data"), err)
}

func TestFindByID_EntityNotFound(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.FindByID("9999")

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, errors.New("failed to fetch data"), err)
}

func setUpTestData(container container.Container) {
	model := model.NewBook("Test1", "123-123-123-1", 1, 1)
	repo := container.GetRepository()
	_, _ = model.Create(repo)
}
