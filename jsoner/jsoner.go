package jsoner

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/utilities"
)

type person struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
}

func (p person) to_json() string {
	u, err := json.Marshal(p)
	utilities.Check(err)
	return string(u)
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
	fmt.Print(carl.to_json())
	utilities.JsonResponse(w, carl)
}
