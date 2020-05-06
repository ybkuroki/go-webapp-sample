package config

import (
	"flag"
	"fmt"

	"github.com/jinzhu/configor"
)

// Config is struct
var Config = struct {
	Database struct {
		Dialect   string `default:"sqlite3"`
		Host      string `default:"book.db"`
		Port      string
		Dbname    string
		Username  string
		Password  string
		Migration bool `default:"true"`
	}
	Extension struct {
		MasterGenerator bool `default:"true"`
	}
}{}

// Load is
func Load() {
	var env = flag.String("env", "develop", "To switch configurations.")
	flag.Parse()
	configor.Load(&Config, "application."+*env+".yml")
	fmt.Println("Loaded this configuration : " + "application." + *env + ".yml")
}
