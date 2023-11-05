package data

import (
	"context"

	"github.com/google/uuid"
	"github.com/ricardoerikson/sgg/pkg/model"
)

// ArticleRepository is the interface to manage articles in the storage layer.
// Every concrete implementation of the storage layer must implement this interface.
type ArticleRepository interface {
	// Create creates a new article in the database
	Create(ctx context.Context, article model.Article) (*model.Article, error)
	// Retrieve retrieves a single article by ID
	Retrieve(ctx context.Context, articleID uuid.UUID) (*model.Article, error)
	// Update updates a single Article
	Update(ctx context.Context, article model.Article) (*model.Article, error)
	// List returns a list of the existing articles
	List(ctx context.Context, withImages string) ([]model.Article, error)
	// Count returns the amount of articles in the database
	Count(ctx context.Context) (int, error)
	// Delete deletes a single article from the database
	Delete(ctx context.Context, articleID uuid.UUID) error
}
