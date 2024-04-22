package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/tjfleming0101/marsairbnb/pkg/config"
	"github.com/tjfleming0101/marsairbnb/pkg/handlers"
	"github.com/tjfleming0101/marsairbnb/pkg/render"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig

const portNumber = ":8080"

var session *scs.SessionManager

// main is the main application function
func main() {

	// Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // In Production this needs to be true

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache:", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Printf("Starting server on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
