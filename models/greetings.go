package models

import (
	"fmt"
	"net/http"
)

func HelloHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func AboutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About page")
}
