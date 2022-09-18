package main

import "net/http"

func (a *app) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/gopher/", a.gopherHandler)

	return mux
}
