package main

import (
	"fmt"
	"net/http"
	"webserver/jsoner"
	"webserver/models"
	"webserver/print"
)

type route struct {
	endpoint string
	function func(w http.ResponseWriter, r *http.Request)
}

var routes = []route{
	{endpoint: "/hello", function: models.HelloHandleFunc},
	{endpoint: "/about", function: models.AboutFunc},
	{endpoint: "/json", function: jsoner.Jsoner},
	{endpoint: "/print", function: print.PrintToFile},
}

func main() {
	for _, v := range routes {
		http.HandleFunc(v.endpoint, v.function)
	}
	port := 8080
	fmt.Printf("Server is listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
