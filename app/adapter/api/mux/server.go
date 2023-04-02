package mux

import (
	"fmt"
	"log"
	"net/http"
	"github.com/anap7/golang-clean-arc-template/app/adapter/api/mux/router"
	//"github.com/anap7/golang-clean-arc-template/app/core/repository"
)

/*type WebServer struct {
	Repository repository.TransactionRepository
}

func NewWebServer() *WebServer {
	return &WebServer{}
}*/
 
func Serve() {
	fmt.Printf("API is Running on port %d", 5000)
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 5000), r))
}


