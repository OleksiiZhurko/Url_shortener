package sever

import (
	"URL_shortener/pkg/config"
	"URL_shortener/pkg/router"
	"log"
	"net/http"
)

// Runs the server
func Run() {
	log.Printf("Starting the server on the port %s\n", config.GetConfigPORT())
	log.Fatal(http.ListenAndServe(config.GetConfigPORT(), router.InitRouter()))
}
