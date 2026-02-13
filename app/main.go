package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fririz/URLShortener/internal/handler"
	"github.com/fririz/URLShortener/internal/middleware"
	"github.com/fririz/URLShortener/internal/repository"
	"github.com/fririz/URLShortener/internal/service"
)

func main() {
	repo, err := repository.NewLinkRepository("./todos.db")
	if err != nil {
		log.Fatal("Error initializing repository: ", err)
	}

	linkSvc, err := service.NewLinkService(repo)
	if err != nil {
		log.Fatal("Error initializing service: ", err)
	}
	linkHdl := handler.NewLinkHandler(linkSvc)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /links", linkHdl.CreateShortLink)
	mux.HandleFunc("GET /{link}", linkHdl.GetFullUrl)

	handlerWithLogging := middleware.LoggingMiddleware(mux)

	port := ":8080"
	fmt.Printf("Server started at http://localhost%s\n", port)

	err = http.ListenAndServe(port, handlerWithLogging)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
