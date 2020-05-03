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

// FindBooksByTitle is
func FindBooksByTitle(title string, page int, size int) *model.Page {
	rep := repository.GetRepository()
	book := model.Book{}
	result, _ := book.FindByTitle(rep, title, page, size)
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

// EditBook is
func EditBook(dto *dto.ChgBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := repository.GetRepository()
		b := model.Book{}
		book, _ := b.FindByID(rep, dto.ID)

		book.Title = dto.Title
		book.Isbn = dto.Isbn

		category := model.Category{}
		book.Category, _ = category.FindByID(rep, dto.CategoryID)

		format := model.Format{}
		book.Format, _ = format.FindByID(rep, dto.FormatID)

		result, _ := book.Save(rep)

		return result, nil
	}

	return nil, errors
}

// DeleteBook is
func DeleteBook(dto *dto.ChgBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := repository.GetRepository()
		b := model.Book{}
		book, _ := b.FindByID(rep, dto.ID)

		result, _ := book.Delete(rep)

		return result, nil
	}

	return nil, errors
}
