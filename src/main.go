package main

import (
	"github.com/charmbracelet/log"
	"github.com/rs/cors"
	"net/http"
	"wennerbeg-wedding/routes"
)

func main() {
	router := routes.Routes()
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	handler := c.Handler(router)
	log.Info("API Server listening on 0.0.0.0:8080")
	http.ListenAndServe(":8080", handler)
}
