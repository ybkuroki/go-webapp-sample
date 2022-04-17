package migration

import (
	"github.com/ybkuroki/go-webapp-sample/container"
	"github.com/ybkuroki/go-webapp-sample/model"
)

// InitMasterData creates the master data used in this application.
func InitMasterData(container container.Container) {
	if container.GetConfig().Extension.MasterGenerator {
		rep := container.GetRepository()

		r := model.NewAuthority("Admin")
		_, _ = r.Create(rep)
		a := model.NewAccountWithPlainPassword("test", "test", r.ID)
		_, _ = a.Create(rep)
		a = model.NewAccountWithPlainPassword("test2", "test2", r.ID)
		_, _ = a.Create(rep)

		c := model.NewCategory("Technical Book")
		_, _ = c.Create(rep)
		c = model.NewCategory("Magazine")
		_, _ = c.Create(rep)
		c = model.NewCategory("Novel")
		_, _ = c.Create(rep)

		f := model.NewFormat("Paper Book")
		_, _ = f.Create(rep)
		f = model.NewFormat("e-Book")
		_, _ = f.Create(rep)
	}
}
