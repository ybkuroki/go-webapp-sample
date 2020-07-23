package model

import (
	"encoding/json"
	"math"

	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Book defines struct of book data.
type Book struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	Title      string    `json:"title"`
	Isbn       string    `json:"isbn"`
	CategoryID uint      `json:"categoryId"`
	Category   *Category `json:"category"`
	FormatID   uint      `json:"formatId"`
	Format     *Format   `json:"format"`
}

// TableName returns the table name of book struct and it is used by gorm.
func (Book) TableName() string {
	return "book"
}

// NewBook is constructor
func NewBook(title string, isbn string, category *Category, format *Format) *Book {
	return &Book{Title: title, Isbn: isbn, Category: category, Format: format}
}

// FindByID returns a book full matched given book's ID.
func (b *Book) FindByID(rep *repository.Repository, id uint) (*Book, error) {
	var book Book
	if error := rep.Preload("Category").Preload("Format").Where("id = ?", id).Find(&book).Error; error != nil {
		return nil, error
	}
	return &book, nil
}

// FindAll returns all books of the book table.
func (b *Book) FindAll(rep *repository.Repository) (*[]Book, error) {
	var books []Book
	if error := rep.Preload("Category").Preload("Format").Find(&books).Error; error != nil {
		return nil, error
	}
	return &books, nil
}

// FindAllByPage returns the page object of all books.
func (b *Book) FindAllByPage(rep *repository.Repository, page int, size int) (*Page, error) {
	var books []Book
	p := createPage(rep, &books, page, size)

	if error := rep.Preload("Category").Preload("Format").Offset(page * p.Size).Limit(size).Find(&books).Error; error != nil {
		return nil, error
	}

	p.Content = &books
	return p, nil
}

// FindByTitle returns the page object of books partially matched given book title.
func (b *Book) FindByTitle(rep *repository.Repository, title string, page int, size int) (*Page, error) {
	var books []Book
	p := createPage(rep, &books, page, size)

	if error := rep.Preload("Category").Preload("Format").Where("title LIKE ?", "%"+title+"%").Offset(page * p.Size).Limit(size).Find(&books).Error; error != nil {
		return nil, error
	}

	p.Content = &books
	return p, nil
}

func createPage(rep *repository.Repository, books *[]Book, page int, size int) *Page {
	p := NewPage()
	p.Page = page
	p.Size = size
	p.NumberOfElements = p.Size

	rep.Preload("Category").Preload("Format").Find(&books).Count(&p.TotalElements)
	p.TotalPages = int(math.Ceil(float64(p.TotalElements) / float64(p.Size)))

	return p
}

// Save persists this book data.
func (b *Book) Save(rep *repository.Repository) (*Book, error) {
	if error := rep.Save(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}

// Update updates this book data.
func (b *Book) Update(rep *repository.Repository) (*Book, error) {
	if error := rep.Update(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}

// Create persists this book data.
func (b *Book) Create(rep *repository.Repository) (*Book, error) {
	if error := rep.Create(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}

// Delete deletes this book data.
func (b *Book) Delete(rep *repository.Repository) (*Book, error) {
	if error := rep.Delete(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}

// ToString is return string of object
func (b *Book) ToString() (string, error) {
	bytes, error := json.Marshal(b)
	return string(bytes), error
}
