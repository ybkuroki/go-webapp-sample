package service

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
)

// CategoryService is a service for managing master data such as format and category.
type CategoryService struct {
	context mycontext.Context
}

// NewCategoryService is constructor.
func NewCategoryService(context mycontext.Context) *CategoryService {
	return &CategoryService{context: context}
}

// FindAllCategories returns the list of all categories.
func (m *CategoryService) FindAllCategories() *[]model.Category {
	rep := m.context.GetRepository()
	category := model.Category{}
	result, err := category.FindAll(rep)
	if err != nil {
		m.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
