package dto

import (
	"encoding/json"

	"github.com/ybkuroki/go-webapp-sample/model"
	"gopkg.in/go-playground/validator.v9"
)

// RegBookDto is struct
type RegBookDto struct {
	Title      string `validate:"required,gte=3,lt=50" json:"title"`
	Isbn       string `validate:"required,gte=10,lt=20" json:"isbn"`
	CategoryID uint   `json:"categoryId"`
	FormatID   uint   `json:"formatId"`
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

// Validate is
func (b *RegBookDto) Validate() map[string]string {
	result := make(map[string]string)
	err := validator.New().Struct(b)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "Title":
					switch errors[i].Tag() {
					case "required", "min", "max":
						result["title"] = "書籍タイトルは、3文字以上50文字以下で入力してください"
					}
				case "Isbn":
					switch errors[i].Tag() {
					case "required", "min", "max":
						result["isbn"] = "ISBNは、10文字以上20文字以下で入力してください"
					}
				}
			}
		}
		return result
	}
	return nil
}

// ToString is return string of object
func (b *RegBookDto) ToString() (string, error) {
	bytes, error := json.Marshal(b)
	return string(bytes), error
}
