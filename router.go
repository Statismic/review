package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
)

func Router() http.Handler {
	r := chi.NewRouter()

	cssHandler := http.FileServer(rice.MustFindBox("public/css").HTTPBox())
	cssHandler = http.StripPrefix("/css/", cssHandler)

	jsHandler := http.FileServer(rice.MustFindBox("public/js").HTTPBox())
	jsHandler = http.StripPrefix("/js/", jsHandler)

	r.Get("/", http.FileServer(rice.MustFindBox("public").HTTPBox()).ServeHTTP)
	r.Get("/css/*", cssHandler.ServeHTTP)
	r.Get("/js/*", jsHandler.ServeHTTP)
	r.Post("/clone", cloneMasterAndPR)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/master/*", masterHandler)
		r.Get("/pr/*", prHandler)
	})
	return r
}
