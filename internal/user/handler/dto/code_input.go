package dto

type CodeInput struct {
	Code string `json:"code" binding:"required"`
}

type RoleInput struct {
	Role string `json:"role" binding:"required"`
}
