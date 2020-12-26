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

		c := model.NewCategory("技術書")
		_, _ = c.Create(rep)
		c = model.NewCategory("雑誌")
		_, _ = c.Create(rep)
		c = model.NewCategory("小説")
		_, _ = c.Create(rep)

		f := model.NewFormat("書籍")
		_, _ = f.Create(rep)
		f = model.NewFormat("電子書籍")
		_, _ = f.Create(rep)
	}
}
