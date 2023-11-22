package entity

import (
	"github.com/google/uuid"
	"time"
)

type FilePath struct {
	Id        uuid.UUID `json:"file_path_id" db:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Mobi      string    `json:"mobi" db:"mobi"`
	Fb2       string    `json:"fb2" db:"fb2"`
	Epub      string    `json:"epub" db:"epub"`
	Docx      string    `json:"docx" db:"docx"`
	CreatedAt time.Time
}
