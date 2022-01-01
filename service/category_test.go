package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestFindAllCategories_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewCategoryService(container)
	result := service.FindAllCategories()

	assert.Len(t, *result, 3)
}
