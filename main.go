package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	// ending request if request is not made to the root.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	message := []byte("Hello from Sinppetbox")
	w.Write(message)
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// extracting value of id parameter
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying snippet with id %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Restricting method to post method only
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// message := []byte("Only POST methods allowed on this route")
		// w.Write(message)
		http.Error(w, "Only POST methods allowed on this route", http.StatusMethodNotAllowed)
		return
	}

	message := []byte("Create a new snippet...")
	w.Write(message)
}

func main() {
	fmt.Println("Hello World!")

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
