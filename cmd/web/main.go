package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", home)
	router.HandleFunc("/gist/create", gistCreate)
	router.HandleFunc("/gist/view", gistView)

	log.Print("starting server at localhost:4000")

	err := http.ListenAndServe("localhost:4000", router)
	log.Fatal(err)
}
