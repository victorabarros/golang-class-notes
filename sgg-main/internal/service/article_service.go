package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ricardoerikson/sgg/pkg/model"
)

type ArticleService interface {
	// CreateArticle creates a new article
	CreateArticle(ctx context.Context, article model.Article) (*model.Article, error)
	// AttachImage attaches an image to an existing article
	AttachImage(ctx context.Context, articleID uuid.UUID, image model.Image) (*model.Article, error)
	// ListArticles return a list of the existing articles. withImages can be [yes,no,*]
	// "yes" - will bring articles with images
	// "no"  - will bring articles without images
	// "*"   - any other value will not apply the filter
	ListArticles(ctx context.Context, withImages string) ([]model.Article, error)
	// RetrieveArticleByID retrieves a single article by ID
	RetrieveArticleByID(ctx context.Context, articleID uuid.UUID) (*model.Article, error)
	// CheckExpired verifies if there are expired articles and automatically deletes
	// them returning the number of deleted articles
	CheckExpired(ctx context.Context) (int, error)
	// CheckImageSize checks if the size of an image is valid
	CheckImageSize(ctx context.Context, url string) error
}
