package entities

import (
	"strings"

	"github.com/google/uuid"
	"github.com/ricardoerikson/sgg/pkg/codes"
	"github.com/ricardoerikson/sgg/pkg/model"
)

// AttachImageRequest holds the request body
type AttachImageRequest struct {
	Path string `json:"path"`
}

// AttachImageParams holds route params
type AttachImageParams struct {
	ArticleID string `schema:"articleId"`
}

// Validate validates attach image params
func (p AttachImageParams) Validate() error {
	if _, err := uuid.Parse(p.ArticleID); err != nil {
		return codes.ErrInvalidArticleID
	}
	return nil
}

// Validate validates an attach image request
func (r AttachImageRequest) Validate() error {
	if strings.TrimSpace(r.Path) == "" {
		return codes.ErrPathIsRequired
	}

	return nil
}

// ToModel converts the request model into an entity model
func (r AttachImageRequest) ToModel() (*model.Image, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	m := model.Image(r)

	return &m, nil
}
