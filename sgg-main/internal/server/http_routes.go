package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ricardoerikson/sgg/pkg/http/middleware"
	"github.com/ricardoerikson/sgg/pkg/http/response"
	"github.com/rs/zerolog/log"
)

type RouteHandler func(*http.Request) response.HTTPResponse

func newRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

func (s *HTTPServer) Wrap(handler RouteHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := handler(r)

		if err := response.WriteResponse(w, resp); err != nil {
			log.Ctx(r.Context()).Error().Err(err).Msg("unexpectedError")
		}
	}
}

func (s *HTTPServer) bindRoutes(r *mux.Router, prefix string) {
	r = r.PathPrefix(prefix).Subrouter()

	r.Use(middleware.ContentType("application/json"))

	articles := r.PathPrefix("/articles").Subrouter()
	articles.HandleFunc("", s.Wrap(s.ah.CreateArticle)).Methods(http.MethodPost)
	articles.HandleFunc("", s.Wrap(s.ah.List)).Methods(http.MethodGet)
	articles.HandleFunc("/{articleId}/images", s.Wrap(s.ah.AttachImage)).Methods(http.MethodPost)
}
