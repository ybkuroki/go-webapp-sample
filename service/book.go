package service

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// FindAllBooks is
func FindAllBooks() *[]model.Book {
	db := repository.GetConnection()
	book := model.Book{}
	result, _ := book.FindAll(db)
	return result
}

// FindAllBooksByPage is
func FindAllBooksByPage(page int, size int) *model.PageDto {
	db := repository.GetConnection()
	book := model.Book{}
	result, _ := book.FindAllByPage(db, page, size)
	return result
}

// RegisterBook is
func RegisterBook(dto *dto.RegBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		db := repository.GetConnection()
		book := dto.Create()

		category := model.Category{}
		book.Category, _ = category.FindByID(db, dto.CategoryID)

		format := model.Format{}
		book.Format, _ = format.FindByID(db, dto.FormatID)

		result, _ := book.Create(db)

		return result, nil
	}

	return nil, errors
}
