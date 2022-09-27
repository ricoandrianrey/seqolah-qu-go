package dto

type RegisterDto struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Role       string `json:"role"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
