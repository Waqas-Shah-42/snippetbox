package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/Waqas-Shah-42/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {

	// setting up custom logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.LUTC)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.LUTC|log.Lshortfile)

	// creating instance of application
	

	infoLog.Println("Starting snippetbox")

	// Getting listening port from commandline
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn","root:example@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()


	// Creating database connection
	db, err := openDb(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	db.Ping()
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
