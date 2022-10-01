package model

import (
	"github.com/moznion/go-optional"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Category defines struct of category data.
type Category struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `validate:"required" json:"name"`
}

// TableName returns the table name of category struct and it is used by gorm.
func (Category) TableName() string {
	return "category_master"
}

// NewCategory is constructor
func NewCategory(name string) *Category {
	return &Category{Name: name}
}

// Exist returns true if a given category exits.
func (c *Category) Exist(rep repository.Repository, id uint) (bool, error) {
	var count int64
	if err := rep.Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// FindByID returns a category full matched given category's ID.
func (c *Category) FindByID(rep repository.Repository, id uint) optional.Option[*Category] {
	var category Category
	if err := rep.Where("id = ?", id).First(&category).Error; err != nil {
		return optional.None[*Category]()
	}
	return optional.Some(&category)
}

// FindAll returns all categories of the category table.
func (c *Category) FindAll(rep repository.Repository) (*[]Category, error) {
	var categories []Category
	if err := rep.Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

// Create persists this category data.
func (c *Category) Create(rep repository.Repository) (*Category, error) {
	if err := rep.Create(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// ToString is return string of object
func (c *Category) ToString() string {
	return toString(c)
}
