package main

import (
	"log"
	"net/http"

	"github.com/andiliewantorosusanto/money-management/pkg/config"
	"github.com/andiliewantorosusanto/money-management/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	handler := routes(&app)

	s := http.Server{
		Addr:    portNumber,
		Handler: handler,
	}

	log.Fatal(s.ListenAndServe())
}
