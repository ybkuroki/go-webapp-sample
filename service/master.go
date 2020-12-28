package service

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
)

// MasterService is a service for managing master data such as format and category.
type MasterService struct {
	context mycontext.Context
}

// NewMasterService is constructor.
func NewMasterService(context mycontext.Context) *MasterService {
	return &MasterService{context: context}
}

// FindAllCategories returns the list of all categories.
func (m *MasterService) FindAllCategories() *[]model.Category {
	rep := m.context.GetRepository()
	category := model.Category{}
	result, err := category.FindAll(rep)
	if err != nil {
		m.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}

// FindAllFormats returns the list of all formats.
func (m *MasterService) FindAllFormats() *[]model.Format {
	rep := m.context.GetRepository()
	format := model.Format{}
	result, err := format.FindAll(rep)
	if err != nil {
		m.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
