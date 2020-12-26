package service

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/model/dto"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// FindAllBooks returns the list of all books.
func FindAllBooks(context mycontext.Context) *[]model.Book {
	rep := context.GetRepository()
	book := model.Book{}
	result, err := book.FindAll(rep)
	if err != nil {
		context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}

// FindAllBooksByPage returns the page object of all books.
func FindAllBooksByPage(context mycontext.Context, page int, size int) *model.Page {
	rep := context.GetRepository()
	book := model.Book{}
	result, err := book.FindAllByPage(rep, page, size)
	if err != nil {
		context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}

// FindBooksByTitle returns the page object of books matched given book title.
func FindBooksByTitle(context mycontext.Context, title string, page int, size int) *model.Page {
	rep := context.GetRepository()
	book := model.Book{}
	result, err := book.FindByTitle(rep, title, page, size)
	if err != nil {
		context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}

// RegisterBook register the given book data.
func RegisterBook(context mycontext.Context, dto *dto.RegBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := context.GetRepository()
		var result *model.Book

		err := rep.Transaction(func(txrep repository.Repository) error {
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
			context.GetLogger().GetZapLogger().Errorf(err.Error())
			return nil, map[string]string{"error": "transaction error"}
		}

		return result, nil
	}

	return nil, errors
}

// EditBook updates the given book data.
func EditBook(context mycontext.Context, dto *dto.ChgBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := context.GetRepository()
		var result *model.Book

		err := rep.Transaction(func(txrep repository.Repository) error {
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
			context.GetLogger().GetZapLogger().Errorf(err.Error())
			return nil, map[string]string{"error": "transaction error"}
		}

		return result, nil
	}

	return nil, errors
}

// DeleteBook deletes the given book data.
func DeleteBook(context mycontext.Context, dto *dto.ChgBookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := context.GetRepository()
		var result *model.Book

		err := rep.Transaction(func(txrep repository.Repository) error {
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
			context.GetLogger().GetZapLogger().Errorf(err.Error())
			return nil, map[string]string{"error": "transaction error"}
		}

		return result, nil
	}

	return nil, errors
}
