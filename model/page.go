package model

// Page defines struct of pagination data.
type Page struct {
	Content          *[]Book `json:"content"`
	Last             bool    `json:"last"`
	TotalElements    int     `json:"totalElements"`
	TotalPages       int     `json:"totalPages"`
	Size             int     `json:"size"`
	Page             int     `json:"page"`
	NumberOfElements int     `json:"numberOfElements"`
}

// NewPage is constructor
func NewPage() *Page {
	return &Page{}
}
