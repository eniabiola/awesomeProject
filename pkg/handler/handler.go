package handler

import (
	"github.com/eniabiola/awesomeProject/pkg/render"
	"net/http"
)

//Home is the home page handler

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}
