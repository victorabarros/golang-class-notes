package core

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type Service interface {
	Start(ctx context.Context, servers ...Server)
}

type service struct {
	name    string
	version string
}

func NewService(name string, version string) Service {
	return &service{
		name:    name,
		version: version,
	}
}

func (s *service) Start(ctx context.Context, servers ...Server) {
	logger := log.Ctx(ctx)
	logger.Info().Msg("starting service.")
	terminationChan := make(chan error, len(servers))

	for _, server := range servers {
		go func(ctx context.Context, server Server) {
			logger.Info().Str("serverName", server.Name()).Msg("starting server.")
			if err := server.Start(ctx); err != nil {
				logger.Warn().Str("serverName", server.Name()).Err(err).Msg("shutting down server")
				terminationChan <- err
			}
		}(ctx, server)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM)

	select {
	case sig := <-signalChan:
		logger.Warn().Str("signal", sig.String()).Msg("shutting down service")
	case err, ok := <-terminationChan:
		if ok {
			logger.Warn().Err(err).Msg("unexpected response from server")
		}
	}
	close(signalChan)

	// graceful shutdown
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(len(servers))

	for _, server := range servers {
		go func(ctx context.Context, server Server) {
			if err := server.Shutdown(ctx); err != nil {
				logger.Info().Str("serverName", server.Name()).Msg("unexpected error during shutdown")
			}

			time.Sleep(time.Second * 2)
			wg.Done()
		}(ctx, server)
	}

	wg.Wait()
	close(terminationChan)
	logger.Info().Msg("service was shutdown.")
}
