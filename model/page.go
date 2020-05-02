package model

// PageDto is
type PageDto struct {
	Content          *[]Book `json:"content"`
	Last             bool    `json:"last"`
	TotalElements    int     `json:"totalElements"`
	TotalPages       int     `json:"totalPages"`
	Size             int     `json:"size"`
	Page             int     `json:"page"`
	NumberOfElements int     `json:"numberOfElements"`
}

// NewPageDto is constructor
func NewPageDto() *PageDto {
	return &PageDto{}
}
