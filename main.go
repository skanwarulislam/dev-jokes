package main

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"log"
	"makeajoke/config"
	"makeajoke/router"
	"net/http"
	"os"
)

func main() {
	key := os.Getenv("SESSION_SECRET")
	maxAge := 86400 * 30 // 30 days
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // to use locally
	store.Options.Secure = false  //For testing
	gothic.Store = store

	var cfg config.Config
	config.ReadFile(&cfg, "config.yml")

	goth.UseProviders(
		google.New(cfg.Providers.Google.ClientId, cfg.Providers.Google.ClientSecret, cfg.Providers.Google.CallbackUrl, "email", "profile"),
		github.New(cfg.Providers.Github.ClientId, cfg.Providers.Github.ClientSecret, cfg.Providers.Github.CallbackUrl),
	)
	routes := router.SetupRouter()
	http.Handle("/", router.SetupRouter())

	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", routes))
}
