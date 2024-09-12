package entity

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	Id          uuid.UUID `json:"author_id" db:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Firstname   string    `json:"firstname" db:"firstname"`
	Lastname    string    `json:"lastname" db:"lastname"`
	ImagePath   string    `json:"image_path" db:"image_path"`
	AboutAuthor string    `json:"about_author" db:"about_author"`
	CreatedAt   time.Time
}
