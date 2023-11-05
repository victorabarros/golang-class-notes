package model

import (
	"time"

	"github.com/google/uuid"
)

// Article holds the article representation
type Article struct {
	ID             uuid.UUID
	Title          string
	Description    string
	ExpirationDate time.Time
	Images         []Image
}

func (a *Article) AttachImage(image Image) {
	if a.Images == nil {
		a.Images = make([]Image, 0)
	}

	a.Images = append(a.Images, image)
}
