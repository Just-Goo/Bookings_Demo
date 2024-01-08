package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Just-Goo/Bookings_Demo/pkg/config"
	"github.com/Just-Goo/Bookings_Demo/pkg/handlers"
	"github.com/Just-Goo/Bookings_Demo/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

// creating a template cache once when 'main' is run and then storing it app wide in 'AppConfig
var app config.AppConfig

var session *scs.SessionManager

func main() {
	
	// change this to 'true' when in production
	app.InProduction = false

	// creating a new session
	session = scs.New()
	session.Lifetime = 24 * time.Hour // This sessions lasts for 24 hours
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session


	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	// Declaring a repository
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println("Listening on port", portNumber) 

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}