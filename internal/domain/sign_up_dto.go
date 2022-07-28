package domain

type SignUpDto struct {
	Login    string `json:"login,omitempty" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required"`
}
