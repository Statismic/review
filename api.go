package main

import "net/http"

type api struct {
	masterDir string
	prDir     string
}

func newAPI(masterDir string, prDir string) *api {
	return &api{
		masterDir: masterDir,
		prDir:     prDir,
	}
}

func (a *api) apiHandle(w http.ResponseWriter, r *http.Request) {
	diff(a.masterDir, a.prDir, "index.html")
	w.WriteHeader(http.StatusOK)
}
