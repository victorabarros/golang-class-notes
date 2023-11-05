package service

import (
	"context"
	"mime"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ricardoerikson/sgg/internal/data"
	"github.com/ricardoerikson/sgg/pkg/codes"
	"github.com/ricardoerikson/sgg/pkg/model"
	"github.com/rs/zerolog/log"
)

// articleService is the concrete implementation of ArticleService
type articleService struct {
	repo data.ArticleRepository
}

// 5 MB
const MaxImageSize = 5 << (10 * 2)

func NewArticleService(repository data.ArticleRepository) ArticleService {
	return &articleService{
		repo: repository,
	}
}

// CreateArticle creates a new article
func (s *articleService) CreateArticle(ctx context.Context, article model.Article) (*model.Article, error) {
	logger := log.Ctx(ctx)

	if article.ExpirationDate.Before(time.Now()) {
		return nil, codes.ErrInvalidExpirationDate
	}

	if len(article.Images) > 3 {
		return nil, codes.ErrImageSizeExceeded
	}

	if len(article.Images) > 0 {
		for _, i := range article.Images {
			if err := s.CheckImageSize(ctx, i.Path); err != nil {
				logger.Info().Err(err).Msg("invalidImageSize")
				return nil, err
			}
		}
	}

	return s.repo.Create(ctx, article)
}

// AttachImage attaches an image to an existing article
func (s *articleService) AttachImage(ctx context.Context, articleID uuid.UUID, image model.Image) (*model.Article, error) {
	logger := log.Ctx(ctx)
	article, err := s.repo.Retrieve(ctx, articleID)
	if err != nil {
		return nil, err
	}

	if len(article.Images) > 2 {
		logger.Info().Err(codes.ErrNumberOfImagesExceeded).Msg("canNotAddMoreImages")
		return nil, codes.ErrNumberOfImagesExceeded
	}

	if err = s.CheckImageSize(ctx, image.Path); err != nil {
		logger.Info().Err(err).Msg("invalidImageSize")
		return nil, err
	}

	article.AttachImage(image)
	updatedArticle, err := s.repo.Update(ctx, *article)
	if err != nil {
		return nil, err
	}

	return updatedArticle, nil
}

// ListArticles return a list of the existing articles
func (s *articleService) ListArticles(ctx context.Context, withImages string) ([]model.Article, error) {
	return s.repo.List(ctx, withImages)
}

// RetrieveArticleByID retrieves a single article by ID
func (s *articleService) RetrieveArticleByID(ctx context.Context, articleID uuid.UUID) (*model.Article, error) {
	return s.repo.Retrieve(ctx, articleID)
}

// CheckExpired verifies if an article is expired and automatically
// deletes it.
func (s *articleService) CheckExpired(ctx context.Context) (int, error) {
	logger := log.Ctx(ctx)

	articles, err := s.repo.List(ctx, "")
	if err != nil {
		logger.Error().Err(err).Msg("unexpectedError")
		return 0, err
	}

	count := 0
	for _, article := range articles {
		if article.ExpirationDate.Before(time.Now()) {
			logger.Info().Str("articleId", article.ID.String()).Msg("articleIsExpired")
			if err := s.repo.Delete(ctx, article.ID); err != nil {
				logger.Error().Err(err).Msg("errorDelitingTheArticle")
			}
			count++
		}
	}
	return count, nil
}

// CheckSize checks if the size of an image is valid
func (s *articleService) CheckImageSize(ctx context.Context, url string) error {
	logger := log.Ctx(ctx)
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Head(url) //nolint:noctx // the method does not accept context
	if err != nil {
		logger.Error().Err(err).Int("code", res.StatusCode).Msg("couldNotReadImage")
		return err
	}
	defer res.Body.Close()

	contentType := res.Header.Get("Content-type")

	isValidType := false
	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			return codes.ErrInvalidImageType
		}

		if strings.HasPrefix(t, "image/") {
			isValidType = true
			break
		}
	}

	if !isValidType {
		return codes.ErrInvalidImageType
	}

	if res.ContentLength > MaxImageSize {
		return codes.ErrImageSizeExceeded
	}

	return nil
}
