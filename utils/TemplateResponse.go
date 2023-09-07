package utils

import (
	"html/template"
	"net/http"
)

func TemplateResponse(filename string, data any, w http.ResponseWriter) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
