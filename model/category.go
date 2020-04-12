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

// Create is
func (c *Category) Create(db *gorm.DB) (*Category, error) {
	if error := db.Create(c).Error; error != nil {
		return nil, error
	}
	return c, nil
}
