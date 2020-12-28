package model

import (
	"encoding/json"

	"github.com/ybkuroki/go-webapp-sample/repository"
)

// Authority defines struct of authority data.
type Authority struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

// TableName returns the table name of authority struct and it is used by gorm.
func (Authority) TableName() string {
	return "authority_master"
}

// NewAuthority is constructor.
func NewAuthority(name string) *Authority {
	return &Authority{Name: name}
}

// Create persists this authority data.
func (a *Authority) Create(rep repository.Repository) (*Authority, error) {
	if error := rep.Create(a).Error; error != nil {
		return nil, error
	}
	return a, nil
}

// ToString is return string of object
func (a *Authority) ToString() (string, error) {
	bytes, error := json.Marshal(a)
	return string(bytes), error
}
