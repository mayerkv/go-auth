package domain

type RefreshDto struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
