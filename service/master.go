package service

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
)

// FindAllCategories returns the list of all categories.
func FindAllCategories(context mycontext.Context) *[]model.Category {
	rep := context.GetRepository()
	category := model.Category{}
	result, err := category.FindAll(rep)
	if err != nil {
		context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}

// FindAllFormats returns the list of all formats.
func FindAllFormats(context mycontext.Context) *[]model.Format {
	rep := context.GetRepository()
	format := model.Format{}
	result, err := format.FindAll(rep)
	if err != nil {
		context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
