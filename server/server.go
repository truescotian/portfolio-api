// Package server provides information required for the sever, plus
// implementation of the server to run.
package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gregpmillr/rest-api/routes"
	"github.com/rs/cors"
)

// Run starts the HTTP listener.
func Run() {
	log.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTP :8080")

	routes.Initialize()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "PATCH", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(routes.RegisterRoutes())

	log.Fatal(http.ListenAndServe(":8080", handler))
}
