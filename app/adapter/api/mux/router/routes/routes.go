package routes

import (
	"net/http"

	"github.com/anap7/golang-clean-arc-template/app/adapter/api/mux/middleware"
	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

//Insert all routes inside the router
func Config(r *mux.Router) *mux.Router {
	routes := transactionRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
	}

	return r
}