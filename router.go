package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/go-chi/chi"
)

func Router() http.Handler {
	r := chi.NewRouter()

	box := rice.MustFindBox("public")
	r.Get("/", http.FileServer(box.HTTPBox()).ServeHTTP)
	r.Post("/clone", cloneMasterAndPR)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/master/*", masterHandler)
		r.Get("/pr/*", prHandler)
	})
	return r
}
