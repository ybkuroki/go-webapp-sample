package config

import (
	"flag"
	"fmt"

	"github.com/jinzhu/configor"
)

// Config is struct
type Config struct {
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
}

const (
	// DEV represents development environment
	DEV = "develop"
	// PRD represents production environment
	PRD = "production"
	// DOC represents docker container
	DOC = "docker"
)

var config *Config
var env *string

// Load is
func Load() {
	env = flag.String("env", "develop", "To switch configurations.")
	flag.Parse()
	config = &Config{}
	configor.Load(config, "application."+*env+".yml")
	fmt.Println("Loaded this configuration : " + "application." + *env + ".yml")
}

// GetConfig is
func GetConfig() *Config {
	return config
}

// GetEnv is
func GetEnv() *string {
	return env
}
