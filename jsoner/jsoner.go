package jsoner

import (
	"fmt"
	"net/http"
)

func Jsoner(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "JSON Returned")
}
