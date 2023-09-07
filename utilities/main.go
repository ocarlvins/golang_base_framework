package utilities

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Quote(s any) string {
	return fmt.Sprintf("\"%s\"", s)
}

func TypeOf(obj any) reflect.Kind {
	return reflect.TypeOf(obj).Kind()
}

func QuoteIfString(value any) string {
	if TypeOf(value) == reflect.String {
		return Quote(value)
	} else {
		return fmt.Sprintf("%v", value)
	}
}

func JsonResponse(w http.ResponseWriter, Object any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	u, err := json.Marshal(Object)
	Check(err)
	_, err = w.Write(u)
	if err != nil {
		panic(err)
	}
}

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
