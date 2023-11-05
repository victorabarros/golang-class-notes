package mem

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/ricardoerikson/sgg/internal/data"
	"github.com/ricardoerikson/sgg/pkg/codes"
	"github.com/ricardoerikson/sgg/pkg/model"
	"github.com/rs/zerolog/log"
)

type articleRepository struct {
	DB   *database
	lock *sync.RWMutex
}

// NewArticleRepository returns an object of the concrete implementation
// of interface ArticleRepository
func NewArticleRepository() data.ArticleRepository {
	return &articleRepository{
		DB: &database{
			storage: make(map[string]model.Article),
		},
		lock: &sync.RWMutex{},
	}
}

// Create creates a new article in the database
func (r *articleRepository) Create(ctx context.Context, a model.Article) (*model.Article, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	a.ID = uuid.New()
	r.DB.storage[a.ID.String()] = a
	log.Ctx(ctx).Info().Msg("articleCreated")

	stored := r.DB.storage[a.ID.String()]
	return &stored, nil
}

// List returns a list of the existing articles
func (r *articleRepository) List(_ context.Context, withImages string) ([]model.Article, error) {
	articles := []model.Article{}
	for _, article := range r.DB.storage {
		if withImages == "yes" && len(article.Images) == 0 {
			continue
		}
		if withImages == "no" && len(article.Images) > 0 {
			continue
		}
		articles = append(articles, article)
	}

	return articles, nil
}

// Retrieve retrieves a single article by ID
func (r *articleRepository) Retrieve(ctx context.Context, articleID uuid.UUID) (*model.Article, error) {
	article, ok := r.DB.storage[articleID.String()]
	if !ok {
		log.Ctx(ctx).Info().Err(codes.ErrArticleNotFound).Msg("articleNotFound")
		return nil, codes.ErrArticleNotFound
	}

	return &article, nil
}

// Update updates a single Article
func (r *articleRepository) Update(ctx context.Context, article model.Article) (*model.Article, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, err := r.Retrieve(ctx, article.ID); err != nil {
		log.Ctx(ctx).Info().Err(err).Msg("couldNotUpdateArticle")
		return nil, err
	}

	r.DB.storage[article.ID.String()] = article

	stored := r.DB.storage[article.ID.String()]
	return &stored, nil
}

// Count returns the amount of articles in the database
func (r *articleRepository) Count(_ context.Context) (int, error) {
	return len(r.DB.storage), nil
}

// Delete deletes an article from the database
func (r *articleRepository) Delete(ctx context.Context, articleID uuid.UUID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, err := r.Retrieve(ctx, articleID); err != nil {
		log.Ctx(ctx).Info().Err(err).Msg("couldNotDeleteArticle")
		return err
	}

	delete(r.DB.storage, articleID.String())

	return nil
}
