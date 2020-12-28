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

// RecordAccount defines struct represents the record of the database.
type RecordAccount struct {
	ID            uint
	Name          string
	Password      string
	AuthorityID   uint
	AuthorityName string
}

const selectAccount = "select a.*, r.id as authority_id, r.name as authority_name " +
	" from account_master a inner join authority_master r on a.authority_id = r.id "

// TableName returns the table name of account struct and it is used by gorm.
func (Account) TableName() string {
	return "account_master"
}

// NewAccount is constructor.
func NewAccount(name string, password string, authorityID uint) *Account {
	return &Account{Name: name, Password: password, AuthorityID: authorityID}
}

// NewAccountWithPlainPassword is constructor. And it is encoded plain text password by using bcrypt.
func NewAccountWithPlainPassword(name string, password string, authorityID uint) *Account {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return &Account{Name: name, Password: string(hashed), AuthorityID: authorityID}
}

// FindByName returns accounts full matched given account name.
func (a *Account) FindByName(rep repository.Repository, name string) (*Account, error) {
	var account *Account

	var rec RecordAccount
	rep.Raw(selectAccount+" where a.name = ?", name).Scan(&rec)
	account = converToAccount(&rec)

	return account, nil
}

// Create persists this account data.
func (a *Account) Create(rep repository.Repository) (*Account, error) {
	if error := rep.Select("name", "password", "authority_id").Create(a).Error; error != nil {
		return nil, error
	}
	return a, nil
}

func converToAccount(rec *RecordAccount) *Account {
	r := &Authority{ID: rec.AuthorityID, Name: rec.AuthorityName}
	return &Account{ID: rec.ID, Name: rec.Name, Password: rec.Password, AuthorityID: rec.AuthorityID, Authority: r}
}

// ToString is return string of object
func (a *Account) ToString() (string, error) {
	bytes, error := json.Marshal(a)
	return string(bytes), error
}
