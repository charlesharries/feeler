package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// home is a wrapper for home routes to return the correct handler depending
// on the request method.
func home(app *application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			homeGet(w, r, app)
		default:
			io.WriteString(w, "{ \"status\": \"method not allowed\" }")
		}
	}
}

// sentiments handles all requests to the /sentiments route and parses
// the right handler function depending on the method.
func sentiments(app *application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			sentimentsPost(w, r, app)
		default:
			io.WriteString(w, "{ \"status\": \"method not allowed\" }")
		}
	}
}

// homeGet handles GET requests to the / route.
func homeGet(w http.ResponseWriter, r *http.Request, app *application) {
	// Render the index page.
	app.render(w, r, "index.page.tmpl")
}

// sentimentPost handles POST requests to the /sentiment route
func sentimentsPost(w http.ResponseWriter, r *http.Request, app *application) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Please submit with Content-Type: application/json", 400)
	}

	req := struct {
		Sentence *string `json:"s"`
	}{}

	json.NewDecoder(r.Body).Decode(&req)

	sent := app.analyser.NewSentiment(*req.Sentence)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sent)
}
