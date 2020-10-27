// Package routes in the starting point for main routes and
// subrouters.
package routes

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gregpmillr/rest-api/controllers"
)

var (
	rootRoute *mux.Router
)

func Initialize() {
	rootRoute = mux.NewRouter()
}

func RegisterRoutes() *negroni.Negroni {
	n := negroni.Classic()
	n.UseHandler(rootRoute)
	rootRoute.HandleFunc("/healthcheck/", controllers.HealthCheck).Methods("GET")
	rootRoute.HandleFunc("/post", controllers.CreatePost).Methods("POST")
	return n
}
