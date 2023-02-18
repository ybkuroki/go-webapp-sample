package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ValidationErrMessageBookTitle string = "Please enter the title with 3 to 50 characters."
	ValidationErrMessageBookISBN  string = "Please enter the ISBN with 10 to 20 characters."
)

func TestValidate_Title2Error(t *testing.T) {
	dto := createBookForTitle2()
	result := dto.Validate()
	assert.Equal(t, ValidationErrMessageBookTitle, result["title"])
}

func TestValidate_Title3Success(t *testing.T) {
	dto := createBookForTitle3()
	result := dto.Validate()
	assert.Empty(t, result)
}

func TestValidate_Title4Success(t *testing.T) {
	dto := createBookForTitle4()
	result := dto.Validate()
	assert.Empty(t, result)
}

func TestValidate_Title49Success(t *testing.T) {
	dto := createBookForTitle49()
	result := dto.Validate()
	assert.Empty(t, result)
}

func TestValidate_Title50Success(t *testing.T) {
	dto := createBookForTitle50()
	result := dto.Validate()
	assert.Empty(t, result)
}

func TestValidate_Title51Error(t *testing.T) {
	dto := createBookForTitle51()
	result := dto.Validate()
	assert.Equal(t, ValidationErrMessageBookTitle, result["title"])
}

func TestValidate_Isbn9Error(t *testing.T) {
	dto := createBookForIsbn9()
	result := dto.Validate()
	assert.Equal(t, ValidationErrMessageBookISBN, result["isbn"])
}

func TestValidate_Isbn10Success(t *testing.T) {
	dto := createBookForIsbn10()
	result := dto.Validate()
	assert.Empty(t, result)
}

func TestValidate_Isbn11Success(t *testing.T) {
	dto := createBookForIsbn11()
	result := dto.Validate()
	assert.Empty(t, result)
}

func TestValidate_Isbn19Success(t *testing.T) {
	dto := createBookForIsbn19()
	result := dto.Validate()
	assert.Empty(t, result)
}

func TestValidate_Isbn20Success(t *testing.T) {
	dto := createBookForIsbn20()
	result := dto.Validate()
	assert.Empty(t, result)
}

func TestValidate_Isbn21Error(t *testing.T) {
	dto := createBookForIsbn21()
	result := dto.Validate()
	assert.Equal(t, ValidationErrMessageBookISBN, result["isbn"])
}

func TestToString(t *testing.T) {
	dto := createBookForTitle4()
	result, _ := dto.ToString()
	assert.Equal(t, "{\"title\":\"Test\",\"isbn\":\"123-123-123-1\",\"categoryId\":1,\"formatId\":1}", result)
}

func createValidationMessages() map[string]string {
	return map[string]string{
		"ValidationErrMessageBookTitle": "Please enter the title with 3 to 50 characters.",
		"ValidationErrMessageBookISBN":  "Please enter the ISBN with 10 to 20 characters."}
}

func createBookForTitle2() *BookDto {
	return &BookDto{
		Title:      "Te",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForTitle3() *BookDto {
	return &BookDto{
		Title:      "Tes",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForTitle4() *BookDto {
	return &BookDto{
		Title:      "Test",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForTitle49() *BookDto {
	return &BookDto{
		Title:      "Test012345Test012345Test012345Test012345Test01234",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForTitle50() *BookDto {
	return &BookDto{
		Title:      "Test012345Test012345Test012345Test012345Test012345",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForTitle51() *BookDto {
	return &BookDto{
		Title:      "Test012345Test012345Test012345Test012345Test012345T",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForIsbn9() *BookDto {
	return &BookDto{
		Title:      "Test",
		Isbn:       "123456789",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForIsbn10() *BookDto {
	return &BookDto{
		Title:      "Test",
		Isbn:       "1234567890",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForIsbn19() *BookDto {
	return &BookDto{
		Title:      "Test",
		Isbn:       "1234567890123456789",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForIsbn20() *BookDto {
	return &BookDto{
		Title:      "Test",
		Isbn:       "12345678901234567890",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForIsbn21() *BookDto {
	return &BookDto{
		Title:      "Test",
		Isbn:       "123456789012345678901",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}

func createBookForIsbn11() *BookDto {
	return &BookDto{
		Title:      "Test",
		Isbn:       "12345678901",
		CategoryID: 1,
		FormatID:   1,
		messages:   createValidationMessages(),
	}
}
