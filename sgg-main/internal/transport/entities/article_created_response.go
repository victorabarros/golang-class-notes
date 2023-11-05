package entities

import (
	"github.com/google/uuid"
	"github.com/ricardoerikson/sgg/pkg/model"
)

type image struct {
	Path string `json:"path"`
}

type ArticleCreatedResponse struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	ExpirationDate string    `json:"expirationDate"`
	Images         []image   `json:"images"`
}

func (r *image) FromModel(m model.Image) {
	r.Path = m.Path
}

func (r *ArticleCreatedResponse) FromModel(m model.Article) {
	r.ID = m.ID
	r.Title = m.Title
	r.Description = m.Description
	r.ExpirationDate = m.ExpirationDate.Format("2006-01-02")
	for _, i := range m.Images {
		r.Images = append(r.Images, image(i))
	}
	if r.Images == nil {
		r.Images = make([]image, 0)
	}
}
