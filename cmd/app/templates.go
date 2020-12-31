package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime/debug"
)

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// We don't want to log helpers.go:14 as the source of the error, but one
	// level up. For that reason we're using log.Logger.Output() to set the
	// frame depth to 2 here.
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// render renders out the given template to the ResponseWriter.
func (app *application) render(w http.ResponseWriter, r *http.Request, name string) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, struct{}{})
	if err != nil {
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}

// newTemplateCache initialises a cache of compiled templates, so we don't have to
// recompile templates on every request.
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		// if err != nil {
		// 	return nil, err
		// }

		cache[name] = ts
	}

	return cache, nil
}
