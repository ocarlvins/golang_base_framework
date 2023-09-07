package views

import (
	"fmt"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Custom 404 Page: Page not found!")
}
