package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `json:"id"  db:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	FirstName  string    `json:"firstname" binding:"required" db:"first_name"`
	LastName   string    `json:"lastname" binding:"required" db:"last_name"`
	Password   string    `json:"password" binding:"required" db:"password"`
	Email      string    `json:"email" binding:"required" db:"email" gorm:"unique"`
	IsVerified bool      `db:"is_verified"`
	CreatedAt  time.Time
}
