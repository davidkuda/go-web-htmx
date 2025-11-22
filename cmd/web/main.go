package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/davidkuda/bellevue/internal/envcfg"
	"github.com/davidkuda/bellevue/internal/models"

	"github.com/coreos/go-oidc/v3/oidc"
	_ "github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/oauth2"
)

type application struct {
	models        models.Models
	templateCache map[string]*template.Template
	OIDC          openIDConnect
}

type openIDConnect struct {
	provider    *oidc.Provider
	verifier    *oidc.IDTokenVerifier
	config      oauth2.Config
	accessToken string
}

var (
	clientID     = os.Getenv("OIDC_CLIENT_ID")
	clientSecret = os.Getenv("OIDC_CLIENT_SECRET")
	issuer       = os.Getenv("OIDC_ISSUER")
	redirectURL  = os.Getenv("OIDC_REDIRECT_URL")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	addr := flag.String("addr", "localhost:8000", "HTTP network address")

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

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		log.Fatal(err)
	}
	app.OIDC.provider = provider

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	app.OIDC.verifier = provider.Verifier(oidcConfig)

	app.OIDC.config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	log.Printf("Starting web server, listening on %s", *addr)
	err = http.ListenAndServe(*addr, app.routes())
	log.Fatal(err)
}
