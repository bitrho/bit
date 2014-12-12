package main

import (
	"log"
	"net/http"

	"github.com/bitrho/bit/router/bit"
	"github.com/gorilla/mux"
)

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	router := mux.NewRouter()
	router.StrictSlash(true)
	router.Handle("/", bit.NewIndex())
	router.Handle("/index/", bit.NewIndex())
	router.Handle("/blog/{date}/{title}/", bit.NewBlog())
	router.Handle("/search/", bit.NewSearch())

	http.Handle("/", router)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Listen error: %v", err)
	}
}
