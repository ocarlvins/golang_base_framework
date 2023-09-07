package views

import (
	"net/http"
	"webserver/utilities"
)

func DefaultPage(w http.ResponseWriter, _ *http.Request) {
	// Define the data to pass to the template
	data := struct {
		Title     string
		Message   string
		FruitList []string
	}{
		Title:     "Hello, Template!",
		Message:   "This is a dynamic message from Go!",
		FruitList: []string{"Apple", "Banana", "Cherry", "Date"},
	}

	utilities.TemplateResponse("templates/index.gohtml", data, w)
}
