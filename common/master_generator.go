package common

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// InitMasterData is
func InitMasterData() {
	db := repository.GetDB()

	db.DropTableIfExists(&model.Book{})
	db.DropTableIfExists(&model.Category{})
	db.DropTableIfExists(&model.Format{})

	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Format{})

	rep := repository.GetRepository()

	c := model.NewCategory("技術書")
	c.Create(rep)
	c = model.NewCategory("雑誌")
	c.Create(rep)
	c = model.NewCategory("小説")
	c.Create(rep)

	f := model.NewFormat("書籍")
	f.Create(rep)
	f = model.NewFormat("電子書籍")
	f.Create(rep)
}
