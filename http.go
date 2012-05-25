package main

import (
	"io"
	"os"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	if path[len(path) - 1] == '/' {
		path += "index.html"
	}

	file, err := os.Open("/www" + path)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
		} else if os.IsPermission(err) {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		io.Copy(w, file)
		file.Close()
	}
}
