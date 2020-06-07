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
		var result *model.Book

		_ = rep.Transaction(func(txrep *repository.Repository) error {
			var err error
			book := dto.Create()

			category := model.Category{}
			if book.Category, err = category.FindByID(txrep, dto.CategoryID); err != nil {
				return err
			}

			format := model.Format{}
			if book.Format, err = format.FindByID(txrep, dto.FormatID); err != nil {
				return err
			}

			if result, err = book.Create(txrep); err != nil {
				return err
			}

			return nil
		})

		return result, nil
	}

	return nil, errors
}

// EditBook is
func EditBook(dto *dto.ChgBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := repository.GetRepository()
		var result *model.Book

		_ = rep.Transaction(func(txrep *repository.Repository) error {
			var err error
			var book *model.Book

			b := model.Book{}
			if book, err = b.FindByID(txrep, dto.ID); err != nil {
				return err
			}

			book.Title = dto.Title
			book.Isbn = dto.Isbn

			category := model.Category{}
			if book.Category, err = category.FindByID(txrep, dto.CategoryID); err != nil {
				return err
			}

			format := model.Format{}
			if book.Format, err = format.FindByID(txrep, dto.FormatID); err != nil {
				return err
			}

			if result, err = book.Save(txrep); err != nil {
				return err
			}

			return nil
		})

		return result, nil
	}

	return nil, errors
}

// DeleteBook is
func DeleteBook(dto *dto.ChgBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := repository.GetRepository()
		var result *model.Book

		_ = rep.Transaction(func(txrep *repository.Repository) error {
			var err error
			var book *model.Book

			b := model.Book{}
			if book, err = b.FindByID(txrep, dto.ID); err != nil {
				return nil
			}

			if result, err = book.Delete(txrep); err != nil {
				return nil
			}

			return nil
		})

		return result, nil
	}

	return nil, errors
}
