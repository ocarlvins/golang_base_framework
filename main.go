package main

import (
	"fmt"
	"net/http"
)

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

type route struct {
	endpoint string
	function func(w http.ResponseWriter, r *http.Request)
}

var routes = []route{
	{endpoint: "/hello", function: helloHandleFunc},
	{endpoint: "/about", function: func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About page")
	}},
}

func main() {
	// http.HandleFunc("/hello", helloHandleFunc)

	// hey := route{
	// 	endpoint: "/hello",
	// 	function: helloHandleFunc,
	// }

	// http.HandleFunc(hey.endpoint, hey.function)

	for _, v := range routes {
		http.HandleFunc(v.endpoint, v.function)
	}
	http.ListenAndServe(":8080", nil)
}
