package model

import (
	"encoding/json"

	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Category is struct
type Category struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `validate:"required" json:"name"`
}

// TableName is
func (Category) TableName() string {
	return "category_master"
}

// NewCategory is constructor
func NewCategory(name string) *Category {
	return &Category{Name: name}
}

// FindByID is
func (c *Category) FindByID(rep *repository.Repository, id uint) (*Category, error) {
	var category Category
	if error := rep.Where("id = ?", id).Find(&category).Error; error != nil {
		return nil, error
	}
	return &category, nil
}

// FindAll is
func (c *Category) FindAll(rep *repository.Repository) (*[]Category, error) {
	var categories []Category
	if error := rep.Find(&categories).Error; error != nil {
		return nil, error
	}
	return &categories, nil
}

// Create is
func (c *Category) Create(rep *repository.Repository) (*Category, error) {
	if error := rep.Create(c).Error; error != nil {
		return nil, error
	}
	return c, nil
}

// ToString is return string of object
func (c *Category) ToString() (string, error) {
	bytes, error := json.Marshal(c)
	return string(bytes), error
}
