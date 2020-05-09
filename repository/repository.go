package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // indirect
	"github.com/ybkuroki/go-webapp-sample/config"
)

// Repository is struct
type Repository struct {
	db *gorm.DB
}

var rep *Repository

const (
	// SQLITE represents SQLite3
	SQLITE = "sqlite3"
	// POSTGRES represents PostgreSQL
	POSTGRES = "postgres"
	// MYSQL represents MySQL
	MYSQL = "mysql"
)

func getConnection(config *config.Config) string {
	if config.Database.Dialect == POSTGRES {
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Dbname, config.Database.Password)
	} else if config.Database.Dialect == MYSQL {
		return fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Dbname)
	}
	return config.Database.Host
}

// InitDB is
func InitDB() {
	db, err := gorm.Open(config.GetConfig().Database.Dialect, getConnection(config.GetConfig()))
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}
	db.LogMode(true)
	rep = &Repository{}
	rep.db = db
}

// GetRepository is
func GetRepository() *Repository {
	return rep
}

// GetDB is
func GetDB() *gorm.DB {
	return rep.db
}

// Find is
func (rep *Repository) Find(out interface{}, where ...interface{}) *gorm.DB {
	return rep.db.Find(out, where...)
}

// Exec is
func (rep *Repository) Exec(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Exec(sql, values...)
}

// First is
func (rep *Repository) First(out interface{}, where ...interface{}) *gorm.DB {
	return rep.db.First(out, where...)
}

// Raw is
func (rep *Repository) Raw(sql string, values ...interface{}) *gorm.DB {
	return rep.db.Raw(sql, values...)
}

// Create is
func (rep *Repository) Create(value interface{}) *gorm.DB {
	return rep.db.Create(value)
}

// Save is
func (rep *Repository) Save(value interface{}) *gorm.DB {
	return rep.db.Save(value)
}

// Update is
func (rep *Repository) Update(value interface{}) *gorm.DB {
	return rep.db.Update(value)
}

// Delete is
func (rep *Repository) Delete(value interface{}) *gorm.DB {
	return rep.db.Delete(value)
}

// Where is
func (rep *Repository) Where(query interface{}, args ...interface{}) *gorm.DB {
	return rep.db.Where(query, args...)
}

// Preload is
func (rep *Repository) Preload(column string, conditions ...interface{}) *gorm.DB {
	return rep.db.Preload(column, conditions...)
}

// Scopes is
func (rep *Repository) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	return rep.db.Scopes(funcs...)
}

// Transaction is
// ref: https://github.com/jinzhu/gorm/blob/master/main.go#L533
func (rep *Repository) Transaction(fc func(tx *Repository) error) (err error) {
	panicked := true
	tx := rep.db.Begin()
	defer func() {
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	txrep := &Repository{}
	txrep.db = tx
	err = fc(txrep)

	if err == nil {
		err = tx.Commit().Error
	}

	panicked = false
	return
}
