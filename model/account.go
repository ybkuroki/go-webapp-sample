package model

import (
	"github.com/ybkuroki/go-webapp-sample/repository"
	"golang.org/x/crypto/bcrypt"
)

// Account is struct
type Account struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Name        string     `json:"name"`
	Password    string     `json:"-"`
	AuthorityID uint       `json:"authority_id"`
	Authority   *Authority `json:"authority"`
}

// TableName is
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

// FindByName is
func (a *Account) FindByName(rep *repository.Repository, name string) (*Account, error) {
	var account Account
	if error := rep.Preload("Authority").Where("name = ?", name).Find(&account).Error; error != nil {
		return nil, error
	}
	return &account, nil
}

// Create persists the data of account object.
func (a *Account) Create(rep *repository.Repository) (*Account, error) {
	if error := rep.Create(a).Error; error != nil {
		return nil, error
	}
	return a, nil
}
