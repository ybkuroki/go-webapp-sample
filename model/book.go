package model

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Book is struct
type Book struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	Title      string    `json:"title"`
	Isbn       string    `json:"isbn"`
	CategoryID uint      `json:"category_id"`
	Category   *Category `json:"category"`
	FormatID   uint      `json:"format_id"`
	Format     *Format   `json:"format"`
}

// NewBook is constructor
func NewBook(title string, isbn string, category *Category, format *Format) *Book {
	return &Book{Title: title, Isbn: isbn, Category: category, Format: format}
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

// SetFormat is setter of Format
func (b *Book) SetFormat(format *Format) {
	b.Format = format
}

// FindByID is
func (b *Book) FindByID(db *gorm.DB, id int) (*Book, error) {
	var book Book
	if error := db.Scopes(repository.Relations(), repository.ByID(id)).Find(&book).Error; error != nil {
		return nil, error
	}
	return &book, nil
}

// FindAll is
func (b *Book) FindAll(db *gorm.DB) (*[]Book, error) {
	var books []Book
	if error := db.Scopes(repository.Relations()).Find(&books).Error; error != nil {
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

// ToString is return string of object
func (b *Book) ToString() (string, error) {
	bytes, error := json.Marshal(b)
	return string(bytes), error
}
