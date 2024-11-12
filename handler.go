package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

type AppMetaData struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSha string `json:"last_commit_sha"`
}

func directory(fileSystem fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		paths := Files(fileSystem)
		for _, i := range paths {
			fmt.Fprintf(w, i+"\n")
		}
	}
}

func Files(fsys fs.FS) (paths []string) {
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		paths = append(paths, p)
		return nil
	})
	return paths
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func metadata(w http.ResponseWriter, _ *http.Request) {
	m := []AppMetaData{
		{
			Version:       Version,
			Description:   Description,
			LastCommitSha: LastCommitSha,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal: %s", err)
	}
	w.Write(jsonResp)

}
