package server

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/ricardoerikson/sgg/internal/cfg"
	"github.com/ricardoerikson/sgg/internal/core"
	httpt "github.com/ricardoerikson/sgg/internal/transport/http"
)

// HTTPServer holds http server setup
type HTTPServer struct {
	API *http.Server

	// handlers
	ah *httpt.ArticleHandler
}

// NewHTTPServer creates a new http server
func NewHTTPServer(ctx context.Context,
	articleHandler *httpt.ArticleHandler,
	config *cfg.Configuration) core.Server {

	h := HTTPServer{
		API: &http.Server{
			Addr:              ":" + strconv.Itoa(config.HTTP.Port),
			BaseContext:       func(l net.Listener) context.Context { return ctx },
			ReadHeaderTimeout: time.Second * 5,
		},
	}

	h.ah = articleHandler

	r := newRouter()
	h.bindRoutes(r, config.Prefix)
	h.API.Handler = r

	return &h
}

func (s *HTTPServer) Name() string {
	return "http-server"
}

func (s *HTTPServer) Start(_ context.Context) error {
	return s.API.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.API.Shutdown(ctx)
}
