package model

import (
	"encoding/json"

	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Format defines struct of format data.
type Format struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `validate:"required" json:"name"`
}

// TableName returns the table name of format struct and it is used by gorm.
func (Format) TableName() string {
	return "format_master"
}

// NewFormat is constructor
func NewFormat(name string) *Format {
	return &Format{Name: name}
}

// FindByID returns a format full matched given format's ID.
func (f *Format) FindByID(rep repository.Repository, id uint) (*Format, error) {
	var format Format
	if error := rep.Where("id = ?", id).First(&format).Error; error != nil {
		return nil, error
	}
	return &format, nil
}

// FindAll returns all formats of the format table.
func (f *Format) FindAll(rep repository.Repository) (*[]Format, error) {
	var formats []Format
	if error := rep.Find(&formats).Error; error != nil {
		return nil, error
	}
	return &formats, nil
}

// Create persists this category data.
func (f *Format) Create(rep repository.Repository) (*Format, error) {
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
