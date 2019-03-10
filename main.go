package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	rice "github.com/GeertJohan/go.rice"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: review <repo_url> <pull_request_id>")
	}

	prID, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal("pull_request_id should be a positive number")
	}

	masterDir, err := repoCloneMaster(os.Args[1])
	if err != nil {
		log.Fatal("error in cleaning master repo")
	}
	defer os.RemoveAll(masterDir)

	prDir, err := repoClonePR(os.Args[1], prID)
	if err != nil {
		log.Fatal("error in cleaning PR repo")
	}
	defer os.RemoveAll(prDir)

	masterHandler := http.FileServer(http.Dir(masterDir))
	prHandler := http.FileServer(http.Dir(prDir))

	http.Handle("/master/", http.StripPrefix("/master/", masterHandler))
	http.Handle("/pr/", http.StripPrefix("/pr/", prHandler))
	http.Handle("/", http.FileServer(rice.MustFindBox("public").HTTPBox()))

	api := newAPI(masterDir, prDir)
	http.HandleFunc("/api/", api.apiHandle)

	go func() {
		<-time.After(100 * time.Millisecond)
		open("http://localhost:8000/")
	}()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
