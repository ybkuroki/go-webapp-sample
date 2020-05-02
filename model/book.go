package model

import (
	"encoding/json"
	"math"

	"github.com/jinzhu/gorm"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Book is struct
type Book struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	Title      string    `json:"title"`
	Isbn       string    `json:"isbn"`
	CategoryID uint      `json:"categoryId"`
	Category   *Category `json:"category"`
	FormatID   uint      `json:"formatId"`
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
func (b *Book) FindByID(db *gorm.DB, id uint) (*Book, error) {
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

// FindAllByPage is
func (b *Book) FindAllByPage(db *gorm.DB, page int, size int) (*PageDto, error) {
	var books []Book

	pagedto := NewPageDto()
	pagedto.Page = page
	pagedto.Size = size
	pagedto.NumberOfElements = pagedto.Size

	db.Scopes(repository.Relations()).Find(&books).Count(&pagedto.TotalElements)
	pagedto.TotalPages = int(math.Ceil(float64(pagedto.TotalElements) / float64(pagedto.Size)))

	if error := db.Scopes(repository.Relations()).Offset(page * pagedto.Size).Limit(size).Find(&books).Error; error != nil {
		return nil, error
	}

	pagedto.Content = &books
	return pagedto, nil
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
