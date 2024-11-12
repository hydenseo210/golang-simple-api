package main

import (
	"net/http"
	"os"
)

var (
	Version       = "1"
	Description   = "pre-interview technical test"
	LastCommitSha = os.Getenv("LASTCOMMITSHA")
)

func main() {

	if os.Getenv("LASTCOMMITSHA") == "" {
		LastCommitSha = "last commit sha not specified"
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)
	http.HandleFunc("/data", metadata)
	http.HandleFunc("/directory", directory(os.DirFS("/")))
	http.ListenAndServe(":8080", nil)
}
