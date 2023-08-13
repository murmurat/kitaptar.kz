package api

import "github.com/google/uuid"

type AuthorRequest struct {
	Id          uuid.UUID `json:"author_id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	ImagePath   string    `json:"image_path"`
	AboutAuthor string    `json:"about_author"`
}
