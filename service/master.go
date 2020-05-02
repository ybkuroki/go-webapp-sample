package service

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// FindAllCategories is
func FindAllCategories() *[]model.Category {
	db := repository.GetConnection()
	category := model.Category{}
	result, _ := category.FindAll(db)
	return result
}

// FindAllFormats is
func FindAllFormats() *[]model.Format {
	db := repository.GetConnection()
	format := model.Format{}
	result, _ := format.FindAll(db)
	return result
}
