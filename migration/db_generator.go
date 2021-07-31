package migration

import (
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
)

// CreateDatabase creates the tables used in this application.
func CreateDatabase(container container.Container) {
	if container.GetConfig().Database.Migration {
		db := container.GetRepository()

		db.DropTableIfExists(&model.Book{})
		db.DropTableIfExists(&model.Category{})
		db.DropTableIfExists(&model.Format{})
		db.DropTableIfExists(&model.Account{})
		db.DropTableIfExists(&model.Authority{})

		db.AutoMigrate(&model.Book{})
		db.AutoMigrate(&model.Category{})
		db.AutoMigrate(&model.Format{})
		db.AutoMigrate(&model.Account{})
		db.AutoMigrate(&model.Authority{})
	}
}
