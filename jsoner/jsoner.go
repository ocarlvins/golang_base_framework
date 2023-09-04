package jsoner

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
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

func Quote(s string) string {
	return fmt.Sprintf("\"%s\"", s)
}

func TypeOf(obj any) reflect.Kind {
	return reflect.TypeOf(obj).Kind()
}

func (p person) insert() string {
	insert_string := fmt.Sprintf("insert into person (first_name) values (")

	if TypeOf(p.FirstName) == reflect.String {
		insert_string += Quote(p.FirstName)
	}

	insert_string += ")"

	return insert_string
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
	fmt.Println(carl.to_json())
	fmt.Println(carl.insert())
	utilities.JsonResponse(w, carl)
}
