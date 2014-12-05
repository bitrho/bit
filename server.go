package main

import (
	"net/http"

	"./middle/logger"
	"./middle/recovery"
	"./middle/static"
	"./router/api"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Handle("/api/{app}/{platform}/{module}/{version}", api.New())

	n := negroni.New()
	n.Use(logger.New())
	n.Use(recovery.New())
	n.Use(static.New(http.Dir("public")))
	n.UseHandler(router)
	n.Run(":3000")
}
