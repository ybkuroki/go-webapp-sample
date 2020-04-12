package model

import (
	"github.com/jinzhu/gorm"
)

// Category is struct
type Category struct {
	gorm.Model
	Name string
}

// NewCategory is constructor
func NewCategory(name string) *Category {
	return &Category{Name: name}
}

// SetName is setter of Name
func (c *Category) SetName(name string) {
	c.Name = name
}

// FindByID is
func (c *Category) FindByID(db *gorm.DB, id int) (*Category, error) {
	var category Category
	if error := db.Find(&category).Error; error != nil {
		return nil, error
	}
	return &category, nil
}

// FindAll is
func (c *Category) FindAll(db *gorm.DB) (*[]Category, error) {
	var categories []Category
	if error := db.Find(&categories).Error; error != nil {
		return nil, error
	}
	return &categories, nil
}

// Create is
func (c *Category) Create(db *gorm.DB) (*Category, error) {
	if error := db.Create(c).Error; error != nil {
		return nil, error
	}
	return c, nil
}
