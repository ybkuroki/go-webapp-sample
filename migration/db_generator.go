package migration

import (
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
)

// CreateDatabase creates the tables used in this application.
func CreateDatabase(container container.Container) {
	if container.GetConfig().Database.Migration {
		db := container.GetRepository()

		_ = db.DropTableIfExists(&model.Book{})
		_ = db.DropTableIfExists(&model.Category{})
		_ = db.DropTableIfExists(&model.Format{})
		_ = db.DropTableIfExists(&model.Account{})
		_ = db.DropTableIfExists(&model.Authority{})

		_ = db.AutoMigrate(&model.Book{})
		_ = db.AutoMigrate(&model.Category{})
		_ = db.AutoMigrate(&model.Format{})
		_ = db.AutoMigrate(&model.Account{})
		_ = db.AutoMigrate(&model.Authority{})
	}
}
