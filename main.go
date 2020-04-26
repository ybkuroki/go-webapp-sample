package main

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"
	"github.com/ybkuroki/go-webapp-sample/router"
)

func main() {
	repository.InitDB()
	db := repository.GetConnection()

	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Format{})

	router := router.Init()
	router.Start(":8080")

	defer db.Close()
}
