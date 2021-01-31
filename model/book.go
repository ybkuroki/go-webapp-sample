package model

import (
	"database/sql"
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

// RecordBook defines struct represents the record of the database.
type RecordBook struct {
	ID           uint
	Title        string
	Isbn         string
	CategoryID   uint
	CategoryName string
	FormatID     uint
	FormatName   string
}

const selectBook = "select b.*, c.id as category_id, c.name as category_name, f.id as format_id, f.name as format_name " +
	"from book b inner join category_master c on c.id = b.category_id inner join format_master f on f.id = b.format_id "

// TableName returns the table name of book struct and it is used by gorm.
func (Book) TableName() string {
	return "book"
}

// NewBook is constructor
func NewBook(title string, isbn string, categoryID uint, formatID uint) *Book {
	return &Book{Title: title, Isbn: isbn, CategoryID: categoryID, FormatID: formatID}
}

// FindByID returns a book full matched given book's ID.
func (b *Book) FindByID(rep repository.Repository, id uint) (*Book, error) {
	var book *Book

	var rec RecordBook
	rep.Raw(selectBook+" where b.id = ?", id).Scan(&rec)
	book = converToBook(&rec)

	return book, nil
}

// FindAll returns all books of the book table.
func (b *Book) FindAll(rep repository.Repository) (*[]Book, error) {
	var books []Book

	var rec RecordBook
	var rows *sql.Rows
	var err error
	if rows, err = rep.Raw(selectBook).Rows(); err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rep.ScanRows(rows, &rec); err != nil {
			return nil, err
		}
		book := converToBook(&rec)
		books = append(books, *book)
	}

	return &books, nil
}

// FindAllByPage returns the page object of all books.
func (b *Book) FindAllByPage(rep repository.Repository, page int, size int) (*Page, error) {
	var books []Book

	var rec RecordBook
	var rows *sql.Rows
	var err error
	if rows, err = rep.Raw(createSQL(selectBook, page, size), size, page*size).Rows(); err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rep.ScanRows(rows, &rec); err != nil {
			return nil, err
		}
		book := converToBook(&rec)
		books = append(books, *book)
	}

	p := createPage(&books, page, size)
	return p, nil
}

// FindByTitle returns the page object of books partially matched given book title.
func (b *Book) FindByTitle(rep repository.Repository, title string, page int, size int) (*Page, error) {
	var books []Book

	var rec RecordBook
	var rows *sql.Rows
	var err error
	if rows, err = rep.Raw(createSQL(selectBook+" where title like ? ", page, size), "%"+title+"%", size, page*size).Rows(); err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rep.ScanRows(rows, &rec); err != nil {
			return nil, err
		}
		book := converToBook(&rec)
		books = append(books, *book)
	}

	p := createPage(&books, page, size)
	return p, nil
}

func createSQL(sql string, page int, size int) string {
	if page > 0 && size > 0 {
		sql += " limit ? offset ? "
	}
	return sql
}

func createPage(books *[]Book, page int, size int) *Page {
	p := NewPage()
	p.Page = page
	p.Size = size
	p.NumberOfElements = p.Size
	p.TotalElements = len(*books)
	p.TotalPages = int(math.Ceil(float64(p.TotalElements) / float64(p.Size)))
	p.Content = books

	return p
}

// Save persists this book data.
func (b *Book) Save(rep repository.Repository) (*Book, error) {
	if error := rep.Save(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}

// Update updates this book data.
func (b *Book) Update(rep repository.Repository) (*Book, error) {
	if error := rep.Model(Book{}).Where("id = ?", b.ID).Select("title", "isbn", "category_id", "format_id").Updates(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}

// Create persists this book data.
func (b *Book) Create(rep repository.Repository) (*Book, error) {
	if error := rep.Select("title", "isbn", "category_id", "format_id").Create(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}

// Delete deletes this book data.
func (b *Book) Delete(rep repository.Repository) (*Book, error) {
	if error := rep.Delete(b).Error; error != nil {
		return nil, error
	}
	return b, nil
}

func converToBook(rec *RecordBook) *Book {
	c := &Category{ID: rec.CategoryID, Name: rec.CategoryName}
	f := &Format{ID: rec.FormatID, Name: rec.FormatName}
	return &Book{ID: rec.ID, Title: rec.Title, Isbn: rec.Isbn, CategoryID: rec.CategoryID, Category: c, FormatID: rec.FormatID, Format: f}
}

// ToString is return string of object
func (b *Book) ToString() (string, error) {
	bytes, error := json.Marshal(b)
	return string(bytes), error
}
