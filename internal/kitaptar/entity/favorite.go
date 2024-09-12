package entity

import (
	"github.com/google/uuid"
	"time"
)

type FavoriteBook struct {
	Id        uuid.UUID `json:"id" db:"id"`
	UserId    string    `json:"user_id" db:"user_id"`
	BookId    string    `json:"book_id" db:"book_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
