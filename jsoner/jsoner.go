package jsoner

import (
	"fmt"
	"net/http"
	"webserver/structs"
	"webserver/utilities"
)

func JsonEndpoint(w http.ResponseWriter, _ *http.Request) {
	carl := structs.Person{
		FirstName:  "FName",
		MiddleName: "Mname",
		LastName:   "LName",
		Age:        45,
	}
	// json.NewEncoder(w).Encode(carl)
	fmt.Println(carl.ToJson())
	fmt.Println(carl.Insert())
	utilities.JsonResponse(w, carl)
}
