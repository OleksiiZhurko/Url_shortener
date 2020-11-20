package router

import (
	"URL_shortener/pkg/router/controller"
	"github.com/julienschmidt/httprouter"
	"log"
)

// Initializing the router
func InitRouter() *httprouter.Router {
	log.Printf("Initializing the router\n")

	handler := controller.NewUrlController()
	router := httprouter.New()
	router.GET("/:shortened", handler.RedirectByShortenLink)
	router.POST("/generate", handler.GenerateShortenLink)

	return router
}