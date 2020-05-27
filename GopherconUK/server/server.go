package server

import "net/http"

func New(mux *http.ServeMux, addr string) *http.Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return srv
}
