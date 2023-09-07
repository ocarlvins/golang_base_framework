package main

import (
	"fmt"
	"net/http"
	"webserver/utils"
	"webserver/views"
)

var routes = []utils.Route{
	{"/", views.DefaultPage},
	{"/hello", views.HelloHandleFunc},
	{"/about", views.AboutFunc},
	{"/json", views.JsonEndpoint},
	{"/print", views.PrintToFile},
}

func main() {
	for _, v := range routes {
		http.HandleFunc(v.Endpoint, v.Function)
	}
	port := 8080
	fmt.Printf("Server is listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return
	}
}
