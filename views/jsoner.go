package views

import (
	"fmt"
	"net/http"
	"webserver/db"
	"webserver/structs"
	"webserver/utils"
)

func JsonEndpoint(w http.ResponseWriter, _ *http.Request) {
	people := []structs.Person{
		{FirstName: "J", MiddleName: "F", LastName: "Kennedy", Age: 75},
		{FirstName: "Barrack", MiddleName: "H", LastName: "Obama", Age: 65},
		{FirstName: "George", MiddleName: "W", Age: 55},
	}
	//carl := structs.Person
	// json.NewEncoder(w).Encode(carl)
	//fmt.Println(carl.ToJson())
	//fmt.Println(carl.Insert())
	for _, v := range people {
		fmt.Println(v.ToJson())
		fmt.Println(db.Insert(v))
	}
	utils.JsonResponse(w, people)
}
