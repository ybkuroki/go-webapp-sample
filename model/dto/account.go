package dto

type LoginDto struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func NewLoginDto() *LoginDto {
	return &LoginDto{}
}
