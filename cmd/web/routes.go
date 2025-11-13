package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/dist/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	standard := alice.New(commonHeaders, logRequest)
	mux.HandleFunc("GET /", app.getHome)

	return standard.Then(mux)
}
