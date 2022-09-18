package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (a *app) gopherHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		if err != nil {
			a.serverError(w)
			return
		}

		g, err := a.gopher.Get(id)
		if err != nil {
			a.serverError(w)
			return
		}

		fmt.Fprintf(w, "%v", *g)
	case http.MethodPost:
		fmt.Fprintf(w, "Creating gopher")
	case http.MethodDelete:
		fmt.Fprintf(w, "Deleting gopher")
	case http.MethodPut:
		fmt.Fprintf(w, "Modifying gopher")
	default:
		w.Header().Set("Allow", "GET, POST, DELETE, PUT")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
