package home

import (
	"log"
	"net/http"
	"time"
)

type Handler struct {
	logger *log.Logger
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello victor\n"))
}

// Logger is a Middleware. I think was a better way to set a middleware.
func (h *Handler) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(start))
		next(w, r)
	}
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
	// Add your more routes here
}

func NewHandler(l *log.Logger) *Handler {
	return &Handler{
		logger: l,
	}
}
