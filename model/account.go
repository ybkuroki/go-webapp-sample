package model

import (
	"encoding/json"

	"github.com/ybkuroki/go-webapp-sample/repository"
	"golang.org/x/crypto/bcrypt"
)

// Account defines struct of account data.
type Account struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Name        string     `json:"name"`
	Password    string     `json:"-"`
	AuthorityID uint       `json:"authority_id"`
	Authority   *Authority `json:"authority"`
}

// TableName returns the table name of account struct and it is used by gorm.
func (Account) TableName() string {
	return "account_master"
}

// NewAccount is constructor.
func NewAccount(name string, password string, authority *Authority) *Account {
	return &Account{Name: name, Password: password, Authority: authority}
}

// NewAccountWithPlainPassword is constructor. And it is encoded plain text password by using bcrypt.
func NewAccountWithPlainPassword(name string, password string, authority *Authority) *Account {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return &Account{Name: name, Password: string(hashed), Authority: authority}
}

// FindByName returns accounts full matched given account name.
func (a *Account) FindByName(rep *repository.Repository, name string) (*Account, error) {
	var account Account
	if error := rep.Preload("Authority").Where("name = ?", name).Find(&account).Error; error != nil {
		return nil, error
	}
	return &account, nil
}

// Create persists this account data.
func (a *Account) Create(rep *repository.Repository) (*Account, error) {
	if error := rep.Create(a).Error; error != nil {
		return nil, error
	}
	return a, nil
}

// ToString is return string of object
func (a *Account) ToString() (string, error) {
	bytes, error := json.Marshal(a)
	return string(bytes), error
}
