package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello someone somewhere"))
}

func gistCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	w.Write([]byte("beautiful gist create point"))
}

func gistView(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || num < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "viewing gist number : %d", num)
}
