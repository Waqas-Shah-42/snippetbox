package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {

	// ending request if request is not made to the root.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",

	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w,"base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	}

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
