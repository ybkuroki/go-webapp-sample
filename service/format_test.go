package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestFindAllFormats_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewFormatService(container)
	result := service.FindAllFormats()

	assert.Len(t, *result, 2)
}
