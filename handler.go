package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/google/uuid"
)

type info struct {
	masterDir string
	prDir     string
}

var handlerMap = newTTLMap(time.Hour * 3)

type repoConfig struct {
	URL string `json:"url"`
	ID  uint64 `json:"id"`
}

func cloneMasterAndPR(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var c repoConfig
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "invalid repoConfig", http.StatusBadRequest)
		return
	}

	masterDir, err := repoCloneMaster(c.URL)
	if err != nil {
		http.Error(w, "failed to clone master repo", http.StatusInternalServerError)
		return
	}

	prDir, err := repoClonePR(c.URL, c.ID)
	if err != nil {
		http.Error(w, "failed to clone PR repo", http.StatusInternalServerError)
		return
	}

	idObj, err := uuid.NewRandom()
	if err != nil {
		http.Error(w, "failed to generate a random uuid4", http.StatusInternalServerError)
		return
	}
	id := idObj.String()

	i := info{
		masterDir: masterDir,
		prDir:     prDir,
	}

	handlerMap.Put(id, &i)

	w.Write([]byte(id))
}

func masterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting master repo")
	id := chi.URLParam(r, "id")
	info := handlerMap.Get(id)

	handler := http.FileServer(http.Dir(info.masterDir))
	handler = http.StripPrefix(fmt.Sprintf("/%s/master", id), handler)
	handler.ServeHTTP(w, r)
}

func prHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	info := handlerMap.Get(id)

	handler := http.FileServer(http.Dir(info.prDir))
	handler = http.StripPrefix(fmt.Sprintf("/%s/pr/", id), handler)
	handler.ServeHTTP(w, r)
}
