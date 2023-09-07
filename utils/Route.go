package utils

import "net/http"

type Route struct {
	Endpoint string
	Function func(w http.ResponseWriter, r *http.Request)
	Methods  []string
}
