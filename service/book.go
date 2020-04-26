package service

import (
	"github.com/jinzhu/gorm"
	"github.com/ybkuroki/go-webapp-sample/model"
)

// FindAllBooks is
func FindAllBooks(db *gorm.DB) *[]model.Book {
	book := model.Book{}
	result, _ := book.FindAll(db)
	return result
}
