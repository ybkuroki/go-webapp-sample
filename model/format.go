package model

import (
	"encoding/json"

	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Format is struct
type Format struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `validate:"required" json:"name"`
}

// TableName is
func (Format) TableName() string {
	return "format_master"
}

// NewFormat is constructor
func NewFormat(name string) *Format {
	return &Format{Name: name}
}

// FindByID is
func (f *Format) FindByID(rep *repository.Repository, id uint) (*Format, error) {
	var format Format
	if error := rep.Where("id = ?", id).Find(&format).Error; error != nil {
		return nil, error
	}
	return &format, nil
}

// FindAll is
func (f *Format) FindAll(rep *repository.Repository) (*[]Format, error) {
	var formats []Format
	if error := rep.Find(&formats).Error; error != nil {
		return nil, error
	}
	return &formats, nil
}

// Create is
func (f *Format) Create(rep *repository.Repository) (*Format, error) {
	if error := rep.Create(f).Error; error != nil {
		return nil, error
	}
	return f, nil
}

// ToString is return string of object
func (f *Format) ToString() (string, error) {
	bytes, error := json.Marshal(f)
	return string(bytes), error
}
