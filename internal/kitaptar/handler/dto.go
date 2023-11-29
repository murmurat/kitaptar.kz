package handler

import (
	"github.com/murat96k/kitaptar.kz/api"
)

type AuthorHandlerDto struct {
	Firstname   string `json:"firstname" `
	Lastname    string `json:"lastname" `
	ImagePath   string `json:"image_path"`
	AboutAuthor string `json:"about_author"`
}

// AuthorRequestBuilder provides an interface for constructing the parts of the user.
type AuthorRequestBuilder interface {
	SetFirstname(firstname string) AuthorRequestBuilder
	SetLastname(lastname string) AuthorRequestBuilder
	SetImagePath(imagePath string) AuthorRequestBuilder
	SetAboutAuthor(aboutAuthor string) AuthorRequestBuilder
	Build() *api.AuthorRequest
}

// NewAuthorBuilder creates a new AuthorRequestBuilder.
func NewAuthorBuilder() AuthorRequestBuilder {
	return &authorBuilder{
		authorRequest: &api.AuthorRequest{}, // Initialize the author attribute
	}
}

type authorBuilder struct {
	authorRequest *api.AuthorRequest
}

func (ab *authorBuilder) SetFirstname(firstname string) AuthorRequestBuilder {
	ab.authorRequest.Firstname = firstname
	return ab
}

func (ab *authorBuilder) SetLastname(lastname string) AuthorRequestBuilder {
	ab.authorRequest.Lastname = lastname
	return ab
}

func (ab *authorBuilder) SetImagePath(imagePath string) AuthorRequestBuilder {
	ab.authorRequest.ImagePath = imagePath
	return ab
}

func (ab *authorBuilder) SetAboutAuthor(aboutAuthor string) AuthorRequestBuilder {
	ab.authorRequest.AboutAuthor = aboutAuthor
	return ab
}

func (ab *authorBuilder) Build() *api.AuthorRequest {
	return ab.authorRequest
}

// Director provides an interface to build authors.
type Director struct {
	builder AuthorRequestBuilder
}

// NewDirector creates a new director for AuthorRequestBuilder.
func NewDirector(builder AuthorRequestBuilder) *Director {
	return &Director{
		builder: builder, // Initialize the builder attribute
	}
}

func (d *Director) ConstructAuthor(firstname, lastname, aboutAuthor, imagePath string) *api.AuthorRequest {
	d.builder.SetFirstname(firstname).
		SetLastname(lastname).
		SetAboutAuthor(aboutAuthor).
		SetImagePath(imagePath)

	return d.builder.Build()
}
