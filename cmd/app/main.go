package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/charlesharries/feeler/pkg/afinn"
	"github.com/charlesharries/feeler/pkg/sentiment"
)

// application holds most of our state that we'll need across the whole app.
type application struct {
	analyser      *sentiment.Analyser
	errorLog      *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	router := http.NewServeMux()

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Read the AFINN data.
	afinn, err := afinn.NewAfinn()
	if err != nil {
		log.Fatal(err)
	}

	// Generate the template cache.
	templateCache, err := newTemplateCache("./web/html")
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		analyser:      sentiment.NewAnalyser(afinn),
		errorLog:      errorLog,
		templateCache: templateCache,
	}

	// Handle static assets.
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Handle routes.
	router.HandleFunc("/", home(app))
	router.HandleFunc("/sentiments", sentiments(app))

	http.ListenAndServe("localhost:3001", router)
}
