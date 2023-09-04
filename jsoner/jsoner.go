package jsoner

import (
	"net/http"
	"webserver/structs"
	"webserver/utilities"
)

func JsonEndpoint(w http.ResponseWriter, _ *http.Request) {
	people := []structs.Person{
		{FirstName: "J", MiddleName: "F", LastName: "Kennedy", Age: 75},
		{FirstName: "Barrack", MiddleName: "H", LastName: "Obama", Age: 65},
		{FirstName: "George", MiddleName: "W", LastName: "Bush", Age: 55},
	}
	//carl := structs.Person
	// json.NewEncoder(w).Encode(carl)
	//fmt.Println(carl.ToJson())
	//fmt.Println(carl.Insert())
	utilities.JsonResponse(w, people)
}
