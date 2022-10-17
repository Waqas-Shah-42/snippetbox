package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// ending request if request is not made to the root.
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",

	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w,err)
		return
	}

	err = ts.ExecuteTemplate(w,"base", nil)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// extracting value of id parameter
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Displaying snippet with id %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Restricting method to post method only
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		//http.Error(w, "Only POST methods allowed on this route", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// message := []byte("Creating a new snippet...")
	// w.Write(message)

	// The hardcoded data will be removed later
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title,content,expires)

	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect user to the relevant page for the snippet
	http.Redirect(w,r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
