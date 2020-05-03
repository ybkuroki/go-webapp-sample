package service

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// FindAllBooks is
func FindAllBooks() *[]model.Book {
	rep := repository.GetRepository()
	book := model.Book{}
	result, _ := book.FindAll(rep)
	return result
}

// FindAllBooksByPage is
func FindAllBooksByPage(page int, size int) *model.Page {
	rep := repository.GetRepository()
	book := model.Book{}
	result, _ := book.FindAllByPage(rep, page, size)
	return result
}

// RegisterBook is
func RegisterBook(dto *dto.RegBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := repository.GetRepository()
		book := dto.Create()

		category := model.Category{}
		book.Category, _ = category.FindByID(rep, dto.CategoryID)

		format := model.Format{}
		book.Format, _ = format.FindByID(rep, dto.FormatID)

		result, _ := book.Create(rep)

		return result, nil
	}

	return nil, errors
}
