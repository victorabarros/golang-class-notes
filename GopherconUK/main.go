package main

import (
	"log"
	"net/http"
	"os"

	"github.com/victorabarros/golang-class-notes/GopherconUK/home"
	"github.com/victorabarros/golang-class-notes/GopherconUK/server"
)

func main() {
	logger := log.New(os.Stdout, "gcUK\t",
		log.LstdFlags|log.Lshortfile)

	h := home.NewHandler(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, ":8081")
	logger.Println("server starting")

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
