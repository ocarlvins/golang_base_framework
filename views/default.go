package views

import (
	"net/http"
	"webserver/utils"
)

func DefaultPage(w http.ResponseWriter, _ *http.Request) {
	//Define the data to pass to the template
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

func ListPage(w http.ResponseWriter, _ *http.Request) {
	data := map[string]interface{}{
		"one":   1,
		"two":   2,
		"three": struct{ Title string }{"Wassup"},
		"four":  4.1234,
	}
	utils.TemplateResponse("templates/list.gohtml", data, w)
}
