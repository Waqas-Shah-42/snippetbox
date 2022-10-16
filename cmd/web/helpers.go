package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Writes error message and stack trace to the errorLog
// Sends 500 error to response user
func (app *application) serverError(w http.ResponseWriter, err error){
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Panicln(trace)

	//so that line number reported is the actual line number and not the line number of the helper file
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int){
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
