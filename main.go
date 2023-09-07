package main

import (
	"fmt"
	"net/http"
	"webserver/jsoner"
	"webserver/models"
	"webserver/print"
	"webserver/views"
)

type route struct {
	endpoint string
	function func(w http.ResponseWriter, r *http.Request)
}

var routes = []route{
	{"/", views.DefaultPage},
	{"/hello", models.HelloHandleFunc},
	{"/about", models.AboutFunc},
	{"/json", jsoner.JsonEndpoint},
	{"/print", print.PrintToFile},
}

func main() {
	for _, v := range routes {
		http.HandleFunc(v.endpoint, v.function)
	}
	port := 8080
	fmt.Printf("Server is listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return
	}
}
