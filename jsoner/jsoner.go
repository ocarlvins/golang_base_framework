package jsoner

import (
	"encoding/json"
	"net/http"
)

type person struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_anme"`
	Age        int    `json:"age"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Jsoner(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "JSON Returned")
	carl := person{
		FirstName:  "FName",
		MiddleName: "Mname",
		LastName:   "LName",
		Age:        45,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(carl)

	u, err := json.Marshal(carl)
	check(err)
	w.Write(u)
}
