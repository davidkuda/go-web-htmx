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

	mux.HandleFunc("GET /{$}", app.getHome)
	mux.HandleFunc("GET /auth", app.getAuth)
	mux.HandleFunc("GET /login", app.oidcLogin)
	mux.HandleFunc("GET /auth/callback", app.oidcCallbackHandler)

	return standard.Then(mux)
}
