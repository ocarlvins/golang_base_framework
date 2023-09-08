package main

import (
	"fmt"
	"net/http"
	"webserver/utils"
	"webserver/views"
)

var GET = []string{"GET"}

var routes = []utils.Route{
	{"/", views.DefaultPage, GET},
	{"/list", views.ListPage, GET},
	{"/hello", views.HelloHandleFunc, GET},
	{"/about", views.AboutFunc, GET},
	{"/json", views.JsonEndpoint, GET},
	{"/print", views.PrintToFile, GET},
	{"/404", views.NotFoundHandler, GET},
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
