package main

import (
	"github.com/anap7/golang-clean-arc-template/app/adapter/api/mux"
	//"github.com/anap7/golang-clean-arc-template/app/handler"
)

func main() {
	// Prepare all environment variables
	//handler.Load()
	// Prepare all routes
	mux.Serve()
}