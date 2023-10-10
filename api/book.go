package api

import "github.com/google/uuid"

type BookRequest struct {
	AuthorId   uuid.UUID `json:"author_id"`
	Annotation string    `json:"annotation"`
	Name       string    `json:"name"`
	Genre      string    `json:"genre"`
	ImagePath  string    `json:"image_path"`
	FilePathId uuid.UUID `json:"file_path_id"`
}
