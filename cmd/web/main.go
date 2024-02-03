package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"

	"github.com/markanator/go-bookings/pkg/config"
	"github.com/markanator/go-bookings/pkg/handlers"
	"github.com/markanator/go-bookings/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	// setup sessions
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Name = "_sid"
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	// set up app config
	app.Session = session

	// set up template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// set up app config
	app.TemplateCache = tc
	app.UseCache = false

	// set up repository
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)
	// set up templates
	render.NewTemplates(&app)

	// create server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// start server
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	err = srv.ListenAndServe()
	log.Fatal(err)
}
