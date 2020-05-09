package migration

import (
	"github.com/ybkuroki/go-webapp-sample/config"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/repository"
)

// InitMasterData is
func InitMasterData(config *config.Config) {
	if config.Extension.MasterGenerator {
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
}
