package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/ricardoerikson/sgg/internal/cfg"
	"github.com/ricardoerikson/sgg/internal/service"
	"github.com/ricardoerikson/sgg/internal/transport/entities"
	"github.com/ricardoerikson/sgg/pkg/http/request"
	"github.com/ricardoerikson/sgg/pkg/http/response"
	"github.com/rs/zerolog/log"
)

type ArticleHandler struct {
	service service.ArticleService
	config  *cfg.Configuration
}

// NewArticleHandler creates a new handler for articles
func NewArticleHandler(service service.ArticleService, config *cfg.Configuration) *ArticleHandler {
	return &ArticleHandler{
		service: service,
		config:  config,
	}
}

// CreateArticle creates a new article
func (h *ArticleHandler) CreateArticle(r *http.Request) response.HTTPResponse {
	reqEntity := new(entities.CreateArticleRequest)

	ctx := r.Context()
	logger := log.Ctx(ctx)

	if err := request.Bind(r, reqEntity); err != nil {
		logger.Info().Err(err).Msg("couldNotBind")
		return response.BadRequest()
	}

	ctx, cancel := context.WithTimeout(ctx, h.config.HTTP.DefaultTimeout)
	defer cancel()
	a, err := reqEntity.ToModel()
	if err != nil {
		return response.ErrorResponse(err)
	}

	newArticle, err := h.service.CreateArticle(ctx, *a)
	if err != nil {
		return response.ErrorResponse(err)
	}

	respEntity := new(entities.ArticleCreatedResponse)
	respEntity.FromModel(*newArticle)

	return response.Created(respEntity)
}

// AttachImage attaches an image to an existing article
func (h *ArticleHandler) AttachImage(r *http.Request) response.HTTPResponse {
	reqEntity := new(entities.AttachImageRequest)

	ctx := r.Context()
	logger := log.Ctx(ctx)
	if err := request.Bind(r, reqEntity); err != nil {
		logger.Info().Err(err).Msg("couldNotBind")
		return response.BadRequest()
	}

	if err := reqEntity.Validate(); err != nil {
		return response.ErrorResponse(err)
	}

	paramsEntity := new(entities.AttachImageParams)
	if err := request.BindParams(r, paramsEntity); err != nil {
		logger.Info().Err(err).Msg("couldNotBindParams")
		return response.BadRequest()
	}

	if err := paramsEntity.Validate(); err != nil {
		return response.ErrorResponse(err)
	}

	ctx, cancel := context.WithTimeout(ctx, h.config.HTTP.DefaultTimeout)
	defer cancel()

	articleID, _ := uuid.Parse(paramsEntity.ArticleID)
	image, err := reqEntity.ToModel()
	if err != nil {
		return response.ErrorResponse(err)
	}

	article, err := h.service.AttachImage(ctx, articleID, *image)
	if err != nil {
		return response.ErrorResponse(err)
	}

	resp := new(entities.ArticleCreatedResponse)
	resp.FromModel(*article)
	return response.Ok(resp)
}

// List returns a list of articles
func (h *ArticleHandler) List(r *http.Request) response.HTTPResponse {
	ctx := r.Context()
	logger := log.Ctx(ctx)

	params := entities.ListArticlesParams{}

	if err := request.BindParams(r, &params); err != nil {
		logger.Info().Msg("couldNotBindParams")
		return response.BadRequest()
	}
	logger.Debug().Interface("params", params).Msg("params")

	ctx, cancel := context.WithTimeout(ctx, h.config.HTTP.DefaultTimeout)
	defer cancel()

	articles, err := h.service.ListArticles(ctx, strings.ToLower(params.WithImages))
	if err != nil {
		logger.Info().Msg("couldNotListArticles")
		return response.BadRequest()
	}

	list := []entities.ArticleCreatedResponse{}
	for _, a := range articles {
		r := entities.ArticleCreatedResponse{}
		r.FromModel(a)
		list = append(list, r)
	}

	return response.Ok(list)
}
