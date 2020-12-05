package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
	"github.com/truescotian/portfolio-api/routes"
)

// Application starting point
func main() {
	client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	routes.Initialize()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "PATCH", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(routes.RegisterRoutes())

	log.Fatal(http.ListenAndServe(":8080", handler))
}
