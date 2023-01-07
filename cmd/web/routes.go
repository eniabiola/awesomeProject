package main

import (
	"github.com/bmizerany/pat"
	"github.com/eniabiola/awesomeProject/pkg/config"
	"github.com/eniabiola/awesomeProject/pkg/handler"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))

	return mux
}
