package service

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// FindAllBooks is
func FindAllBooks() *[]model.Book {
	db := repository.GetConnection()
	book := model.Book{}
	result, _ := book.FindAll(db)
	return result
}
