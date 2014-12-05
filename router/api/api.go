package api

import (
	"log"
	"net/http"
)

type api struct {
	count int
}

func New() *api {
	return &api{}
}

func (a *api) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Log status
	log.Printf(req.URL.Path)
}
