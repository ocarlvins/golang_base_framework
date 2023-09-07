package views

import (
	"net/http"
	"webserver/utilities"
)

func DefaultPage(w http.ResponseWriter, _ *http.Request) {
	// Define the data to pass to the template
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Hello, Template!",
		Message: "This is a dynamic message from Go!",
	}

	utilities.TemplateResponse("templates/index.html", data, w)
}
