package controller

import (
	"URL_shortener/pkg/generators"
	"URL_shortener/pkg/router/response"
	"URL_shortener/pkg/router/response/failed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type UrlController struct {
	links *response.URLs
}

// Returns Url controller
func NewUrlController() *UrlController {
	return &UrlController{links: response.NewURLs()}
}

// Returns a shortened and entered link with code 201 if a request is right
// otherwise returns error with code 400
//
// Type JSON
// POST 127.0.0.1:PORT/generate/{params}
func (uc UrlController) GenerateShortenLink(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	link := retrievePostURL(request)

	if link != "" {
		uc.links.SetParams(link, generators.GenerateRandomString())

		writeResponseInWriter(
			writer,
			http.StatusCreated,
			uc.links.ConvertToJSON())

		log.Printf("A new key was generated for %s\n", link)
		log.Printf("Response with %d code\n", http.StatusCreated)
	} else {
		writeResponseInWriter(
			writer,
			http.StatusBadRequest,
			failed.ProvideBadRequest(fmt.Sprintf("Error in %s", request.Form)))

		log.Printf("User input error in the form %s\n", request.Form)
		log.Printf("Response with %d code\n", http.StatusBadRequest)
	}
}

// Redirects to the specified link with 302 code if it exists
// otherwise returns error with 404
//
// Type JSON
// GET 127.0.0.1:PORT/{link}
func (uc UrlController) RedirectByShortenLink(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	link := ps.ByName("shortened") // retrieve hash

	if uc.links.CompareByShorten(link) {
		http.Redirect(writer, request, uc.links.GetFullLink(), http.StatusFound)

		log.Printf("User has been redirected to %s\n", uc.links.GetFullLink())
		log.Printf("Response with %d code\n", http.StatusFound)
	} else {
		writeResponseInWriter(
			writer,
			http.StatusNotFound,
			failed.ProvideNotFound(link))

		log.Printf("%s does not exist\n", link)
		log.Printf("Response with %d code\n", http.StatusNotFound)
	}
}

// Recording response
func writeResponseInWriter(writer http.ResponseWriter, code int, message []byte) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write(message)
}

// Retrieves url param from POST request
func retrievePostURL(request *http.Request) string {
	if err := request.ParseForm(); err != nil {
		log.Fatalf("Something wrong with %s", request.Body)
		return ""
	}

	if len(request.Form) != 1 {
		return ""
	}

	return request.Form.Get("url")
}
