package api

type AuthorRequest struct {
	Firstname   string `json:"firstname" `
	Lastname    string `json:"lastname" `
	ImagePath   string `json:"image_path"`
	AboutAuthor string `json:"about_author"`
}
