package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `json:"id" db:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	FirstName string    `json:"firstname" db:"first_name"`
	LastName  string    `json:"lastname" db:"last_name"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email" gorm:"unique"`
	CreatedAt time.Time
}
