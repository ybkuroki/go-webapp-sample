package migration

import (
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/mycontext"
)

// InitMasterData creates the master data used in this application.
func InitMasterData(context mycontext.Context) {
	if context.GetConfig().Extension.MasterGenerator {
		rep := context.GetRepository()

		r := model.NewAuthority("Admin")
		_, _ = r.Create(rep)
		a := model.NewAccountWithPlainPassword("test", "test", r.ID)
		_, _ = a.Create(rep)

		c := model.NewCategory("Technical book")
		_, _ = c.Create(rep)
		c = model.NewCategory("Magazine")
		_, _ = c.Create(rep)
		c = model.NewCategory("Novel")
		_, _ = c.Create(rep)

		f := model.NewFormat("Book")
		_, _ = f.Create(rep)
		f = model.NewFormat("E-Book")
		_, _ = f.Create(rep)
	}
}
