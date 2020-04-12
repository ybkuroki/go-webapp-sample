package main

import (
	"fmt"

	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// go get -u github.com/jinzhu/gorm
// go get -u github.com/mattn/go-sqlite3
// go get -u github.com/ybkuroki/go-webapp-sample
func main() {
	db := initDB()

	c := NewCategory("technical")
	c.Create(db)

	book := NewBook("test", "123-123-1", c)
	book.Create(db)

	result, _ := book.FindByID(db, 1)
	fmt.Println(result.Title, result.Isbn, result.Category.Name)

	category := NewCategory("magazine")
	category.Create(db)

	defer db.Close()
}

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Category{})
	return db
}
