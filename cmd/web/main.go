package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	mux := http.NewServeMux()

	// Serve static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// remove /static prefix from path and server the file at /static/
	mux.Handle("/static/", http.StripPrefix("/static",fileServer))

	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
