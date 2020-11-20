package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUrlController_Redirect(t *testing.T) {
	handler := NewUrlController()
	router := httprouter.New()

	router.GET("/:shortened", handler.RedirectByShortenLink)

	handler.links.SetParams("https://developers.google.com/", "hash1d")

	tested := []struct { // data to test
		url string
		code int
	}{
		{url: "/test1d", code: http.StatusNotFound},
		{url: "/hash1", code: http.StatusNotFound},
		{url: "/hash1d", code: http.StatusFound},
	}

	for one := 0; one < len(tested); one++ {
		req, _ := http.NewRequest("GET", tested[one].url, nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != tested[one].code {
			t.Errorf("Wrong status")
		}
	}

	req, _ := http.NewRequest("GET", "/test1d", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Wrong status")
	}
}

func TestUrlController_Generate(t *testing.T) {
	handler := NewUrlController()
	router := httprouter.New()

	router.POST("/generate", handler.GenerateShortenLink)

	tested := []struct { // data to test
		form url.Values
		code int
	}{
		{form: url.Values{"url": []string{"https://translate.google.com/"}},
			code: http.StatusCreated},
		{form: url.Values{"urf": []string{"https://developers.google.com/"}},
			code: http.StatusBadRequest},
		{form: url.Values{"url": []string{"https://developers.google.com/"},
							"urf": []string{"https://developers.google.com/"}},
			code: http.StatusBadRequest},
	}

	for one := 0; one < len(tested); one++ {
		req, _ := http.NewRequest("POST", "/generate", strings.NewReader(tested[one].form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != tested[one].code {
			t.Errorf("Wrong status")
		}
	}
}
