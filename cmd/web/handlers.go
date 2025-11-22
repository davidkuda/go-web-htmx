package main

import (
	"errors"
	"fmt"
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

// GET /auth
func (app *application) getAuth(w http.ResponseWriter, r *http.Request) {
	t := app.newTemplateData(r)
	fmt.Println(t.User)
	app.render(w, r, http.StatusOK, "auth.tmpl.html", &t)
}
