package api

type FilePathRequest struct {
	Mobi string `json:"mobi"`
	Fb2  string `json:"fb2"`
	Epub string `json:"epub"`
	Docx string `json:"docx"`
}
