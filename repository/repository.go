package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // indirect
)

var db *gorm.DB
var err error

// InitDB is
func InitDB() {
	db, err = gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
}

// GetConnection is
func GetConnection() *gorm.DB {
	return db
}

// Relations is
func Relations() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Category").Preload("Format")
	}
}

// ByID is
func ByID(id int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

// ByName is
func ByName(name string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name LIKE ?", "%"+name+"%")
	}
}

// ByTitle is
func ByTitle(title string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("title LIKE ?", "%"+title+"%")
	}
}
