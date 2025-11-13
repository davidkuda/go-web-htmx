package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/davidkuda/bellevue/internal/envcfg"
	"github.com/davidkuda/bellevue/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	models models.Models

	templateCache     map[string]*template.Template
	templateCacheHTMX map[string]*template.Template
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	addr := flag.String("addr", ":8875", "HTTP network address")

	app := &application{}

	db, err := envcfg.DB()
	if err != nil {
		log.Fatalf("could not open DB: %v\n", err)
	}
	defer db.Close()

	app.models = models.New(db)

	app.templateCache, err = newTemplateCache(true)
	if err != nil {
		log.Fatalf("could not initialise templateCache: %v\n", err)
	}

	app.templateCacheHTMX, err = newTemplateCache(false)
	if err != nil {
		log.Fatalf("could not initialise templateCache: %v\n", err)
	}

	log.Printf("Starting web server, listening on %s", *addr)
	err = http.ListenAndServe(*addr, app.routes())
	log.Fatal(err)
}
