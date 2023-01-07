package main

import (
	"fmt"
	"github.com/eniabiola/awesomeProject/pkg/config"
	"github.com/eniabiola/awesomeProject/pkg/handler"
	"github.com/eniabiola/awesomeProject/pkg/render"
	"log"
	"net/http"
)

const portNumber string = ":8080"

// main is the main application function
func main() {

	var app config.AppConfig

	//get the template cache from the app config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)
	/*http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)*/

	fmt.Printf("Starting application on %s", portNumber)

	/*_ = http.ListenAndServe(
		portNumber,
		nil,
	)*/

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
