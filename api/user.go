package api

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

//type RegisterRequest struct {
//	Email     string `json:"email" db:"email" binding:"required"`
//	FirstName string `json:"first_name" db:"first_name" binding:"required"`
//	LastName  string `json:"last_name" db:"last_name" binding:"required"`
//	Password  string `json:"password" db:"password" binding:"required"`
//}
