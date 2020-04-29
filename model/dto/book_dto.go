package dto

import "github.com/ybkuroki/go-webapp-sample/model"

// RegBookDto is struct
type RegBookDto struct {
	Title      string
	Isbn       string
	CategoryID uint
	FormatID   uint
}

// NewRegBookDto is
func NewRegBookDto() *RegBookDto {
	return &RegBookDto{}
}

// Create is
func (b *RegBookDto) Create() *model.Book {
	c := model.NewCategory("")
	f := model.NewFormat("")
	return model.NewBook(b.Title, b.Isbn, c, f)
}
