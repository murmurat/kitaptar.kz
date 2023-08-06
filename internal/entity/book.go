package entity

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	Id         uuid.UUID `json:"book_id" db:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	AuthorId   uuid.UUID `json:"author_id" db:"author_id" gorm:"type:uuid"`
	Author     Author    `gorm:"foreignKey:author_id"`
	Annotation string    `json:"annotation" db:"annotation"`
	Name       string    `json:"name" db:"name"`
	Genre      string    `json:"genre" db:"genre"`
	ImagePath  string    `json:"image_path" db:"image_path"`
	CreatedAt  time.Time
}
