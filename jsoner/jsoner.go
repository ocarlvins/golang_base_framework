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

func JsonResponse(w http.ResponseWriter, Object any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	u, err := json.Marshal(Object)
	check(err)
	w.Write(u)
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
	JsonResponse(w, carl)

}
