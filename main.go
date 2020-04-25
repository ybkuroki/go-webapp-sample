package main

import (
	"github.com/ybkuroki/go-webapp-sample/controller"
	"github.com/ybkuroki/go-webapp-sample/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// go get -u github.com/jinzhu/gorm
// go get -u github.com/mattn/go-sqlite3
// go get -u github.com/ybkuroki/go-webapp-sample
func main() {
	db := initDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/book/list", controller.BookList(db))

	e.Start(":8080")

	//c := model.NewCategory("technical")
	//c.Create(db)

	//f := model.NewFormat("paper")
	//f.Create(db)

	//book := model.NewBook("test", "123-123-1", c, f)
	//book.Create(db)

	//result, _ := book.FindByID(db, 1)
	//fmt.Println(result.Title, result.Isbn, result.Category.Name, result.Format.Name)

	//category := model.NewCategory("magazine")
	//category.Create(db)

	defer db.Close()
}

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Format{})
	return db
}
