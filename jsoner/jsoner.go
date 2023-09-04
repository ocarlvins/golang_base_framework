package jsoner

import (
	"net/http"
	"webserver/utilities"
)

type person struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_anme"`
	Age        int    `json:"age"`
}

func Jsoner(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "JSON Returned")
	carl := person{
		FirstName:  "FName",
		MiddleName: "Mname",
		LastName:   "LName",
		Age:        45,
	}
	// json.NewEncoder(w).Encode(carl)
	utilities.JsonResponse(w, carl)
}
