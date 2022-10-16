package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// setting up custom logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

	// creating instance of application
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	
	infoLog.Println("Starting snippetbox")

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()

	// Serve static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// remove /static prefix from path and server the file at /static/
	mux.Handle("/static/", http.StripPrefix("/static",fileServer))

	mux.HandleFunc("/", app.home)

	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on port :%v", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
