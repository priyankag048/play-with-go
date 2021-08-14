package main

import (
	"log"
	"net/http"
	"time"

	"template_render/pkg/config"
	"template_render/pkg/handlers"
	"template_render/pkg/renderer"

	"github.com/alexedwards/scs/v2"
)

const port = ":8001"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	renderer.NewTemplates(&app)

	handler := getHandler(&app)
	log.Fatal(http.ListenAndServe(port, handler))
}
