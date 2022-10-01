package model

import (
	"database/sql"
	"errors"
	"math"

	"github.com/moznion/go-optional"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/util"
	"gorm.io/gorm"
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

const (
	selectBook = "select b.id as id, b.title as title, b.isbn as isbn, " +
		"c.id as category_id, c.name as category_name, f.id as format_id, f.name as format_name " +
		"from book b inner join category_master c on c.id = b.category_id inner join format_master f on f.id = b.format_id "
	findByID    = " where b.id = ?"
	findByTitle = " where title like ? "
)

// TableName returns the table name of book struct and it is used by gorm.
func (Book) TableName() string {
	return "book"
}

// NewBook is constructor
func NewBook(title string, isbn string, categoryID uint, formatID uint) *Book {
	return &Book{Title: title, Isbn: isbn, CategoryID: categoryID, FormatID: formatID}
}

// FindByID returns a book full matched given book's ID.
func (b *Book) FindByID(rep repository.Repository, id uint) optional.Option[*Book] {
	var rec RecordBook
	args := []interface{}{id}

	createRaw(rep, selectBook+findByID, "", "", args).Scan(&rec)
	return convertToBook(&rec)
}

// FindAll returns all books of the book table.
func (b *Book) FindAll(rep repository.Repository) (*[]Book, error) {
	var books []Book
	var err error

	if books, err = findRows(rep, selectBook, "", "", []interface{}{}); err != nil {
		return nil, err
	}
	return &books, nil
}

// FindAllByPage returns the page object of all books.
func (b *Book) FindAllByPage(rep repository.Repository, page string, size string) (*Page, error) {
	var books []Book
	var err error

	if books, err = findRows(rep, selectBook, page, size, []interface{}{}); err != nil {
		return nil, err
	}
	p := createPage(&books, page, size)
	return p, nil
}

// FindByTitle returns the page object of books partially matched given book title.
func (b *Book) FindByTitle(rep repository.Repository, title string, page string, size string) (*Page, error) {
	var books []Book
	var err error
	args := []interface{}{"%" + title + "%"}

	if books, err = findRows(rep, selectBook+findByTitle, page, size, args); err != nil {
		return nil, err
	}
	p := createPage(&books, page, size)
	return p, nil
}

func findRows(rep repository.Repository, sqlquery string, page string, size string, args []interface{}) ([]Book, error) {
	var books []Book

	var rec RecordBook
	var rows *sql.Rows
	var err error

	if rows, err = createRaw(rep, sqlquery, page, size, args).Rows(); err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rep.ScanRows(rows, &rec); err != nil {
			return nil, err
		}

		opt := convertToBook(&rec)
		if opt.IsNone() {
			return nil, errors.New("failed to fetch data")
		}
		book, _ := opt.Take()
		books = append(books, *book)
	}
	return books, nil
}

func createRaw(rep repository.Repository, sql string, pageNum string, pageSize string, args []interface{}) *gorm.DB {
	if util.IsNumeric(pageNum) && util.IsNumeric(pageSize) {
		page := util.ConvertToInt(pageNum)
		size := util.ConvertToInt(pageSize)
		args = append(args, size)
		args = append(args, page*size)
		sql += " limit ? offset ? "
	}
	if len(args) > 0 {
		return rep.Raw(sql, args...)
	}
	return rep.Raw(sql)
}

func createPage(books *[]Book, page string, size string) *Page {
	p := NewPage()
	p.Page = util.ConvertToInt(page)
	p.Size = util.ConvertToInt(size)
	p.NumberOfElements = p.Size
	p.TotalElements = len(*books)
	if p.TotalPages = int(math.Ceil(float64(p.TotalElements) / float64(p.Size))); p.TotalPages < 0 {
		p.TotalPages = 0
	}
	p.Content = books

	return p
}

// Save persists this book data.
func (b *Book) Save(rep repository.Repository) (*Book, error) {
	if err := rep.Save(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

// Update updates this book data.
func (b *Book) Update(rep repository.Repository) (*Book, error) {
	if err := rep.Model(Book{}).Where("id = ?", b.ID).Select("title", "isbn", "category_id", "format_id").Updates(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

// Create persists this book data.
func (b *Book) Create(rep repository.Repository) (*Book, error) {
	if err := rep.Select("title", "isbn", "category_id", "format_id").Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

// Delete deletes this book data.
func (b *Book) Delete(rep repository.Repository) (*Book, error) {
	if err := rep.Delete(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func convertToBook(rec *RecordBook) optional.Option[*Book] {
	if rec.ID == 0 {
		return optional.None[*Book]()
	}
	c := &Category{ID: rec.CategoryID, Name: rec.CategoryName}
	f := &Format{ID: rec.FormatID, Name: rec.FormatName}
	return optional.Some(
		&Book{ID: rec.ID, Title: rec.Title, Isbn: rec.Isbn, CategoryID: rec.CategoryID, Category: c, FormatID: rec.FormatID, Format: f})
}

// ToString is return string of object
func (b *Book) ToString() string {
	return toString(b)
}
