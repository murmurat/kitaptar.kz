package dto

type RefreshInput struct {
	Token string `json:"token" binding:"required"`
}
