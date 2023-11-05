package server

import (
	"context"

	"github.com/ricardoerikson/sgg/internal/cfg"
	"github.com/ricardoerikson/sgg/internal/core"
	"github.com/ricardoerikson/sgg/internal/service"
	"github.com/robfig/cron"
	"github.com/rs/zerolog/log"
)

type cronServer struct {
	articlesService service.ArticleService
	c               *cron.Cron
	ctx             context.Context
	config          *cfg.Configuration
}

func NewCronServer(ctx context.Context, service service.ArticleService, config *cfg.Configuration) core.Server {
	return &cronServer{
		c:               cron.New(),
		articlesService: service,
		ctx:             ctx,
		config:          config,
	}
}

func (s *cronServer) Name() string {
	return "cron-server"
}

func (s *cronServer) Start(_ context.Context) error {
	logger := log.Ctx(s.ctx)
	err := s.c.AddFunc(s.config.Cron.DeleteArticles, func() {
		logger.Info().Msg("running delete expired articles job")
		count, err := s.articlesService.CheckExpired(s.ctx)
		if err != nil {
			logger.Error().Err(err).Msg("errorDeletingArticles")
			return
		}

		if count > 0 {
			logger.Info().Int("count", count).Msg("articlesDeleted")
			return
		}

		logger.Info().Msg("noArticlesDeleted")
	})

	if err != nil {
		return err
	}

	s.c.Start()
	return nil
}

func (s *cronServer) Shutdown(_ context.Context) error {
	s.c.Stop()
	return nil
}
