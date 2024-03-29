package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"code.google.com/p/gorilla/sessions"
)

const (
	sessionKeyLoc = ".sessions.key"
)

var store *sessions.CookieStore

func init() {
	key, err := ioutil.ReadFile(sessionKeyLoc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No session keys available, check %s.\n", sessionKeyLoc)
		os.Exit(1)
	}
	store = sessions.NewCookieStore(key)
}

func GetSession(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "IND_SESSID")
	return session
}
