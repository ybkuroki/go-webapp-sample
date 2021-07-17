package dto

import (
	"encoding/json"

	"github.com/ybkuroki/go-webapp-sample/model"
	"gopkg.in/go-playground/validator.v9"
)

const (
	required string = "required"
	max      string = "max"
	min      string = "min"
)

// BookDto defines a data transfer object for book.
type BookDto struct {
	Title      string `validate:"required,min=3,max=50" json:"title"`
	Isbn       string `validate:"required,min=10,max=20" json:"isbn"`
	CategoryID uint   `json:"categoryId"`
	FormatID   uint   `json:"formatId"`
}

// NewBookDto is constructor.
func NewBookDto() *BookDto {
	return &BookDto{}
}

// Create creates a book model from this DTO.
func (b *BookDto) Create() *model.Book {
	return model.NewBook(b.Title, b.Isbn, b.CategoryID, b.FormatID)
}

// Validate performs validation check for the each item.
func (b *BookDto) Validate() map[string]string {
	return validateDto(b)
}

func validateDto(b interface{}) map[string]string {
	result := make(map[string]string)
	err := validator.New().Struct(b)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "ID":
					switch errors[i].Tag() {
					case required:
						result["id"] = "書籍IDが存在しません"
					}
				case "Title":
					switch errors[i].Tag() {
					case required, min, max:
						result["title"] = "書籍タイトルは、3文字以上50文字以下で入力してください"
					}
				case "Isbn":
					switch errors[i].Tag() {
					case required, min, max:
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
func (b *BookDto) ToString() (string, error) {
	bytes, error := json.Marshal(b)
	return string(bytes), error
}
