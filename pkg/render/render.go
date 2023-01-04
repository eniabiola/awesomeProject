package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders template using html
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("there is an error ", err)
		return
	}
}
