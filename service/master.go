package service

import (
	"github.com/ybkuroki/go-webapp-sample/logger"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// FindAllCategories returns the list of all categories.
func FindAllCategories() *[]model.Category {
	rep := repository.GetRepository()
	category := model.Category{}
	result, err := category.FindAll(rep)
	if err != nil {
		logger.GetEchoLogger().Error(err.Error)
		return nil
	}
	return result
}

// FindAllFormats returns the list of all formats.
func FindAllFormats() *[]model.Format {
	rep := repository.GetRepository()
	format := model.Format{}
	result, err := format.FindAll(rep)
	if err != nil {
		logger.GetEchoLogger().Error(err.Error)
		return nil
	}
	return result
}
