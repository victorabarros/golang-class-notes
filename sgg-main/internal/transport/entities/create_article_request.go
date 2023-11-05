package entities

import (
	"strings"
	"time"

	"github.com/ricardoerikson/sgg/pkg/codes"
	"github.com/ricardoerikson/sgg/pkg/model"
)

type CreateArticleRequest struct {
	Title          string               `json:"title"`
	Description    string               `json:"description"`
	ExpirationDate string               `json:"expirationDate"`
	Images         []AttachImageRequest `json:"images"`
}

// Validate validates a create article request
func (r CreateArticleRequest) Validate() error {
	if strings.TrimSpace(r.Title) == "" {
		return codes.ErrTitleIsRequired
	}

	if strings.TrimSpace(r.Description) == "" {
		return codes.ErrDescriptionIsRequired
	}

	if len(strings.TrimSpace(r.Description)) > 4000 {
		return codes.ErrDescriptionIsTooLong
	}

	if strings.TrimSpace(r.ExpirationDate) == "" {
		return codes.ErrExpirationDateIsRequired
	}

	if len(r.Images) > 3 {
		return codes.ErrNumberOfImagesExceeded
	}

	_, err := time.Parse("2006-01-02", r.ExpirationDate)
	if err != nil {
		return codes.ErrInvalidExpirationDate
	}

	return nil
}

// ToModel converts a request entity into an internal model
func (r CreateArticleRequest) ToModel() (*model.Article, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	expirationDate, _ := time.Parse("2006-01-02", r.ExpirationDate)

	images := make([]model.Image, 0)
	for _, reqImage := range r.Images {
		i, err := reqImage.ToModel()
		if err != nil {
			return nil, err
		}

		images = append(images, *i)
	}

	return &model.Article{
		Title:          r.Title,
		Description:    r.Description,
		ExpirationDate: expirationDate,
		Images:         images,
	}, nil
}
