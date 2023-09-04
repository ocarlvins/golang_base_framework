package main

import (
	"net/http"
	"webserver/jsoner"
	"webserver/models"
)

type route struct {
	endpoint string
	function func(w http.ResponseWriter, r *http.Request)
}

var routes = []route{
	{endpoint: "/hello", function: models.HelloHandleFunc},
	{endpoint: "/about", function: models.AboutFunc},
	{endpoint: "/json", function: jsoner.Jsoner},
}

func main() {
	for _, v := range routes {
		http.HandleFunc(v.endpoint, v.function)
	}
	http.ListenAndServe(":8080", nil)
}
