package model

import (
	"github.com/ybkuroki/go-webapp-sample/repository"

	"github.com/jinzhu/gorm"
)

// Book is struct
type Book struct {
	gorm.Model
	Title      string
	Isbn       string
	CategoryID int
	Category   *Category
}

// NewBook is constructor
func NewBook(title string, isbn string, category *Category) *Book {
	return &Book{Title: title, Isbn: isbn, Category: category}
}

// SetTitle is setter of Title
func (b *Book) SetTitle(title string) {
	b.Title = title
}

// SetIsbn is setter of Isbn
func (b *Book) SetIsbn(isbn string) {
	b.Isbn = isbn
}

// SetCategory is setter of Category
func (b *Book) SetCategory(category *Category) {
	b.Category = category
}

// FindByID is
func (b *Book) FindByID(db *gorm.DB, id int) (*Book, error) {
	var book Book
	if error := db.Scopes(Relations(), ByID(id)).Find(&book).Error; error != nil {
		return nil, error
	}
	return &book, nil
}

// FindAll is
func (b *Book) FindAll(db *gorm.DB) (*[]Book, error) {
	var books []Book
	if error := db.Scopes(Relations()).Find(&books).Error; error != nil {
		return nil, error
	}
	return &books, nil
}

// Save is
func (b *Book) Save(db *gorm.DB) (*Book, error) {
	if error := db.Save(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}

// Update is
func (b *Book) Update(db *gorm.DB) (*Book, error) {
	if error := db.Update(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}

// Create is
func (b *Book) Create(db *gorm.DB) (*Book, error) {
	if error := db.Create(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}
