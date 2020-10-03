package service

import (
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// FindAllBooks returns the list of all books.
func FindAllBooks() *[]model.Book {
	rep := repository.GetRepository()
	book := model.Book{}
	result, err := book.FindAll(rep)
	if err != nil {
		logger.GetEchoLogger().Error(err.Error)
		return nil
	}
	return result
}

// FindAllBooksByPage returns the page object of all books.
func FindAllBooksByPage(page int, size int) *model.Page {
	rep := repository.GetRepository()
	book := model.Book{}
	result, err := book.FindAllByPage(rep, page, size)
	if err != nil {
		logger.GetEchoLogger().Error(err.Error)
		return nil
	}
	return result
}

// FindBooksByTitle returns the page object of books matched given book title.
func FindBooksByTitle(title string, page int, size int) *model.Page {
	rep := repository.GetRepository()
	book := model.Book{}
	result, err := book.FindByTitle(rep, title, page, size)
	if err != nil {
		logger.GetEchoLogger().Error(err.Error)
		return nil
	}
	return result
}

// RegisterBook register the given book data.
func RegisterBook(dto *dto.RegBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := repository.GetRepository()
		var result *model.Book

		err := rep.Transaction(func(txrep *repository.Repository) error {
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

		if err != nil {
			logger.GetEchoLogger().Error(err)
			return nil, map[string]string{"error": "transaction error"}
		}

		return result, nil
	}

	return nil, errors
}

// EditBook updates the given book data.
func EditBook(dto *dto.ChgBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := repository.GetRepository()
		var result *model.Book

		err := rep.Transaction(func(txrep *repository.Repository) error {
			var err error
			var book *model.Book

			b := model.Book{}
			if book, err = b.FindByID(txrep, dto.ID); err != nil {
				return err
			}

			book.Title = dto.Title
			book.Isbn = dto.Isbn
			book.CategoryID = dto.CategoryID
			book.FormatID = dto.FormatID

			category := model.Category{}
			if book.Category, err = category.FindByID(txrep, dto.CategoryID); err != nil {
				return err
			}

			format := model.Format{}
			if book.Format, err = format.FindByID(txrep, dto.FormatID); err != nil {
				return err
			}

			if result, err = book.Update(txrep); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			logger.GetEchoLogger().Error(err)
			return nil, map[string]string{"error": "transaction error"}
		}

		return result, nil
	}

	return nil, errors
}

// DeleteBook deletes the given book data.
func DeleteBook(dto *dto.ChgBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := repository.GetRepository()
		var result *model.Book

		err := rep.Transaction(func(txrep *repository.Repository) error {
			var err error
			var book *model.Book

			b := model.Book{}
			if book, err = b.FindByID(txrep, dto.ID); err != nil {
				return err
			}

			if result, err = book.Delete(txrep); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			logger.GetEchoLogger().Error(err)
			return nil, map[string]string{"error": "transaction error"}
		}

		return result, nil
	}

	return nil, errors
}
