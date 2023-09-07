package views

import (
	"net/http"
	"webserver/utils"
)

func DefaultPage(w http.ResponseWriter, _ *http.Request) {
	// Define the data to pass to the template
	data := struct {
		Title     string
		Message   string
		FruitList []string
	}{
		"Hello, Template!",
		"This is a dynamic message from Go!",
		[]string{"Apple", "Banana", "Cherry", "Date"},
	}

	utils.TemplateResponse("templates/index.gohtml", data, w)
}
