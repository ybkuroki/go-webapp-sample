package model

import "github.com/ybkuroki/go-webapp-sample/repository"

// Authority is struct
type Authority struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

// TableName is
func (Authority) TableName() string {
	return "authority_master"
}

// NewAuthority is constructor.
func NewAuthority(name string) *Authority {
	return &Authority{Name: name}
}

// Create persists the data of account object.
func (a *Authority) Create(rep *repository.Repository) (*Authority, error) {
	if error := rep.Create(a).Error; error != nil {
		return nil, error
	}
	return a, nil
}
