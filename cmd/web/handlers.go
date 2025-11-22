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
	app.render(w, r, http.StatusOK, "home.tmpl.html", &t)
}
