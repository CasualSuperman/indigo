package main

import (
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/", rootHandler)
//	http.Handle("/", http.FileServer(http.Dir("/www")))
	http.ListenAndServe(":8080", nil)
}
