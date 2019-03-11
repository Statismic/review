package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Router() http.Handler {
	r := chi.NewRouter()

	cssHandler := http.FileServer(http.Dir("public/css"))
	cssHandler = http.StripPrefix("/css/", cssHandler)

	jsHandler := http.FileServer(http.Dir("public/js"))
	jsHandler = http.StripPrefix("/js/", jsHandler)

	svgHandler := http.FileServer(http.Dir("public/svg"))
	svgHandler = http.StripPrefix("/svg/", svgHandler)

	r.Get("/", http.FileServer(http.Dir("public")).ServeHTTP)
	r.Get("/css/*", cssHandler.ServeHTTP)
	r.Get("/js/*", jsHandler.ServeHTTP)
	r.Get("/svg/*", svgHandler.ServeHTTP)
	r.Post("/clone", cloneMasterAndPR)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/master/*", masterHandler)
		r.Get("/pr/*", prHandler)
	})
	return r
}
