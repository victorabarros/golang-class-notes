package main

import (
	"context"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ricardoerikson/sgg/internal/cfg"
	"github.com/ricardoerikson/sgg/internal/core"
	"github.com/ricardoerikson/sgg/internal/data/mem"
	"github.com/ricardoerikson/sgg/internal/server"
	"github.com/ricardoerikson/sgg/internal/service"
	"github.com/ricardoerikson/sgg/internal/transport/grpc"
	"github.com/ricardoerikson/sgg/internal/transport/http"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	config := new(cfg.Configuration)
	if err := cleanenv.ReadEnv(config); err != nil {
		log.Error().Err(err).Msg("could not read environment variables")
		return
	}

	serviceName := "articles-service"

	ctx = log.With().
		Caller().
		Str("service", serviceName).
		Str("version", Version).Logger().
		WithContext(ctx)

	articleRepository := mem.NewArticleRepository()
	articleService := service.NewArticleService(articleRepository)
	articleHandlerHTTP := http.NewArticleHandler(articleService, config)
	articleHandlerGRPC := grpc.NewArticleHandler(articleService, config)

	httpServer := server.NewHTTPServer(ctx, articleHandlerHTTP, config)
	grpcServer := server.NewGRPCServer(ctx, articleHandlerGRPC, config)
	cronServer := server.NewCronServer(ctx, articleService, config)

	svc := core.NewService(serviceName, Version)
	svc.Start(ctx, httpServer, grpcServer, cronServer)
}
