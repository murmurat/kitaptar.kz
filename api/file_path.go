package api

import "github.com/google/uuid"

type FilePathRequest struct {
	Id   uuid.UUID `json:"file_path_id"`
	Mobi string    `json:"mobi"`
	Fb2  string    `json:"fb2"`
	Epub string    `json:"epub"`
	Docx string    `json:"docx"`
}
