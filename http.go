package main

import (
	"fmt"
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
		session := GetSession(r)
		fmt.Println(session)
		session.Values["hello"] = "world"
		session.Save(r, w)
		io.Copy(w, file)
		file.Close()
	}
}
