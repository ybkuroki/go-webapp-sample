package common

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// InitMasterData is
func InitMasterData() {
	db := repository.GetConnection()

	db.DropTableIfExists(&model.Book{})
	db.DropTableIfExists(&model.Category{})
	db.DropTableIfExists(&model.Format{})

	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Format{})

	c := model.NewCategory("技術書")
	c.Create(db)
	c = model.NewCategory("雑誌")
	c.Create(db)
	c = model.NewCategory("小説")
	c.Create(db)

	f := model.NewFormat("書籍")
	f.Create(db)
	f = model.NewFormat("電子書籍")
	f.Create(db)
}
