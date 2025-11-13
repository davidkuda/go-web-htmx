package main

import (
	"errors"
	"net/http"
)

var (
	FieldError = errors.New("FieldError")
)

// GET /
func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
	t := app.newTemplateData(r)

	// TODO: move isHTMX logic into app.render
	isHTMX := r.Header.Get("HX-Request") == "true"
	if isHTMX {
		app.renderHTMXPartial(w, r, http.StatusOK, "home.tmpl.html", &t)
	} else {
		app.render(w, r, http.StatusOK, "home.tmpl.html", &t)
	}
}
