package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/ricardoerikson/sgg/internal/cfg"
	"github.com/ricardoerikson/sgg/internal/service"
	"github.com/ricardoerikson/sgg/internal/transport/pb"
	"github.com/rs/zerolog/log"
)

type ArticleHandler struct {
	*pb.UnimplementedArticlesServer
	service service.ArticleService
	config  *cfg.Configuration
}

func NewArticleHandler(service service.ArticleService, config *cfg.Configuration) *ArticleHandler {
	return &ArticleHandler{
		service: service,
		config:  config,
	}
}

func (h *ArticleHandler) Create(ctx context.Context, r *pb.CreateArticleRequest) (*pb.Article, error) {
	req := ConvertCreateArticleRequestFromProtoToEntity(r)

	ctx, cancel := context.WithTimeout(ctx, h.config.HTTP.DefaultTimeout)
	defer cancel()

	a, err := req.ToModel()
	if err != nil {
		return nil, err
	}

	article, err := h.service.CreateArticle(ctx, *a)
	if err != nil {
		return nil, err
	}

	res := ConvertArticleFromModelToProto(*article)
	return res, nil
}

func (h *ArticleHandler) List(ctx context.Context, r *pb.ListArticlesRequest) (*pb.ArticlesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.config.HTTP.DefaultTimeout)
	defer cancel()

	articles, err := h.service.ListArticles(ctx, r.WithImages)
	if err != nil {
		return nil, err
	}

	res := &pb.ArticlesResponse{}
	for _, article := range articles {
		protoArticle := ConvertArticleFromModelToProto(article)
		res.Articles = append(res.Articles, protoArticle)
	}

	return res, nil
}

// AttachImage attaches an image to an existing
func (h *ArticleHandler) AttachImage(ctx context.Context, r *pb.AttachImageRequest) (*pb.Article, error) {
	req := ConvertAttachImageRequestFromProtoToEntity(r)

	logger := log.Ctx(ctx)
	id, err := uuid.Parse(r.ArticleId)
	if err != nil {
		logger.Info().Err(err).Msg("invalidRequest")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, h.config.HTTP.DefaultTimeout)
	defer cancel()

	i, err := req.ToModel()
	if err != nil {
		return nil, err
	}

	article, err := h.service.AttachImage(ctx, id, *i)
	if err != nil {
		return nil, err
	}

	protoArticle := ConvertArticleFromModelToProto(*article)

	return protoArticle, nil
}
