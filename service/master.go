package service

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// FindAllCategories is
func FindAllCategories() *[]model.Category {
	rep := repository.GetRepository()
	category := model.Category{}
	result, _ := category.FindAll(rep)
	return result
}

// FindAllFormats is
func FindAllFormats() *[]model.Format {
	rep := repository.GetRepository()
	format := model.Format{}
	result, _ := format.FindAll(rep)
	return result
}
