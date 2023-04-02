package router

import (
	"net/http"
	"github.com/anap7/golang-clean-arc-template/app/adapter/api/mux/router/routes"
	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

//It will generate a new router with all routes configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	//Call a config that prepare all routes
	return routes.Config(r)
}

