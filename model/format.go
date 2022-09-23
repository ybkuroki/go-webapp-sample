package model

import (
	"github.com/moznion/go-optional"
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
func (f *Format) FindByID(rep repository.Repository, id uint) optional.Option[*Format] {
	var format Format
	if err := rep.Where("id = ?", id).First(&format).Error; err != nil {
		return optional.None[*Format]()
	}
	return optional.Some(&format)
}

// FindAll returns all formats of the format table.
func (f *Format) FindAll(rep repository.Repository) (*[]Format, error) {
	var formats []Format
	if err := rep.Find(&formats).Error; err != nil {
		return nil, err
	}
	return &formats, nil
}

// Create persists this category data.
func (f *Format) Create(rep repository.Repository) (*Format, error) {
	if err := rep.Create(f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

// ToString is return string of object
func (f *Format) ToString() string {
	return toString(f)
}
