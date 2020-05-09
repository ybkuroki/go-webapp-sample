package migration

import (
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// CreateDatabase is
func CreateDatabase(config *config.Config) {
	if config.Database.Migration {
		db := repository.GetDB()

		db.DropTableIfExists(&model.Book{})
		db.DropTableIfExists(&model.Category{})
		db.DropTableIfExists(&model.Format{})

		db.AutoMigrate(&model.Book{})
		db.AutoMigrate(&model.Category{})
		db.AutoMigrate(&model.Format{})
	}
}
