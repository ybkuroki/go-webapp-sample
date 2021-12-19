package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
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

func TestFindAllBooks_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.FindAllBooks()

	assert.Equal(t, 2, len(*result))
	assert.Equal(t, nil, err)
}

func TestFindAllBooksByPage_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.FindAllBooksByPage("0", "5")

	assert.Equal(t, 2, result.TotalElements)
	assert.Equal(t, 1, result.TotalPages)
	assert.Equal(t, 0, result.Page)
	assert.Equal(t, 5, result.Size)
	assert.Equal(t, 2, len(*result.Content))
	assert.Equal(t, nil, err)
}

func TestFindBooksByTitle_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.FindBooksByTitle("1", "0", "5")

	assert.Equal(t, 1, result.TotalElements)
	assert.Equal(t, 1, result.TotalPages)
	assert.Equal(t, 0, result.Page)
	assert.Equal(t, 5, result.Size)
	assert.Equal(t, 1, len(*result.Content))
	assert.Equal(t, nil, err)
}

func TestCreateBook_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewBookService(container)
	result, err := service.CreateBook(createBookForCreate())

	entity := &model.Book{}
	data, _ := entity.FindByID(container.GetRepository(), 1)

	assert.Equal(t, data, result)
	assert.Equal(t, map[string]string(map[string]string(nil)), err)
}

func TestCreateBook_IsbnLess_ValidationError(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewBookService(container)
	result, err := service.CreateBook(createBookForIsbnLessValidationError())

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, "ISBNは、10文字以上20文字以下で入力してください", err["isbn"])
}

func TestCreateBook_IsbnMore_ValidationError(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewBookService(container)
	result, err := service.CreateBook(createBookForIsbnMoreValidationError())

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, "ISBNは、10文字以上20文字以下で入力してください", err["isbn"])
}

func TestCreateBook_TitleLess_ValidationError(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewBookService(container)
	result, err := service.CreateBook(createBookForTitleLessValidationError())

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, "書籍タイトルは、3文字以上50文字以下で入力してください", err["title"])
}

func TestCreateBook_TitleMore_ValidationError(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewBookService(container)
	result, err := service.CreateBook(createBookForTitleMoreValidationError())

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, "書籍タイトルは、3文字以上50文字以下で入力してください", err["title"])
}

func TestCreateBook_NotCategory(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewBookService(container)
	result, err := service.CreateBook(createBookForNotCategory())

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, "Failed to the registration", err["error"])
}

func TestCreateBook_NotFormat(t *testing.T) {
	container := test.PrepareForServiceTest()

	service := NewBookService(container)
	result, err := service.CreateBook(createBookForNotFormat())

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, "Failed to the registration", err["error"])
}

func TestUpdateBook_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.UpdateBook(createBookForCreate(), "1")

	entity := &model.Book{}
	data, _ := entity.FindByID(container.GetRepository(), 1)

	assert.Equal(t, data, result)
	assert.Equal(t, map[string]string(map[string]string(nil)), err)
}

func TestUpdateBook_NotEntity(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.UpdateBook(createBookForNotCategory(), "99")

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, "Failed to the update", err["error"])
}

func TestUpdateBook_NotCategory(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.UpdateBook(createBookForNotCategory(), "1")

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, "Failed to the update", err["error"])
}

func TestUpdateBook_NotFormat(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.UpdateBook(createBookForNotFormat(), "1")

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, "Failed to the update", err["error"])
}

func TestDeleteBook_Success(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	entity := &model.Book{}
	data, _ := entity.FindByID(container.GetRepository(), 1)

	service := NewBookService(container)
	result, err := service.DeleteBook("1")

	assert.Equal(t, data, result)
	assert.Equal(t, map[string]string(map[string]string(nil)), err)
}

func TestDeleteBook_Error(t *testing.T) {
	container := test.PrepareForServiceTest()

	setUpTestData(container)

	service := NewBookService(container)
	result, err := service.DeleteBook("99")

	assert.Equal(t, (*model.Book)(nil), result)
	assert.Equal(t, "Failed to the delete", err["error"])
}

func setUpTestData(container container.Container) {
	entity := model.NewBook("Test1", "123-123-123-1", 1, 1)
	repo := container.GetRepository()
	_, _ = entity.Create(repo)

	entity = model.NewBook("Test2", "123-123-123-2", 2, 2)
	_, _ = entity.Create(repo)
}

func createBookForCreate() *dto.BookDto {
	return &dto.BookDto{
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
	}
}

func createBookForIsbnLessValidationError() *dto.BookDto {
	return &dto.BookDto{
		Title:      "Test",
		Isbn:       "123",
		CategoryID: 1,
		FormatID:   1,
	}
}

func createBookForIsbnMoreValidationError() *dto.BookDto {
	return &dto.BookDto{
		Title:      "Test",
		Isbn:       "123-123-123-123-123-123",
		CategoryID: 1,
		FormatID:   1,
	}
}

func createBookForTitleLessValidationError() *dto.BookDto {
	return &dto.BookDto{
		Title:      "T",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
	}
}

func createBookForTitleMoreValidationError() *dto.BookDto {
	return &dto.BookDto{
		Title:      "TestTestTestTestTestTestTestTestTestTestTestTestTest",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   1,
	}
}

func createBookForNotCategory() *dto.BookDto {
	return &dto.BookDto{
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: 99,
		FormatID:   1,
	}
}

func createBookForNotFormat() *dto.BookDto {
	return &dto.BookDto{
		Title:      "Test1",
		Isbn:       "123-123-123-1",
		CategoryID: 1,
		FormatID:   99,
	}
}
