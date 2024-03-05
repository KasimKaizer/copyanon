package main

import "net/http"

func (app *application) router() *http.ServeMux {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./ui/static"))
	router.Handle("/static/", http.StripPrefix("/static", fs))

	router.HandleFunc("/", app.home)
	router.HandleFunc("/gist/create", app.gistCreate)
	router.HandleFunc("/gist/view", app.gistView)
	return router
}
