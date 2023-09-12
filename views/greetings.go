package views

import (
	"fmt"
	"net/http"
)

func HelloHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
	//w.Write([]byte("Hello, world"))
	if r.Method == http.MethodGet {
		if r.URL.Path == "/foo" {

		}

	}
}

func AboutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About page")
}
