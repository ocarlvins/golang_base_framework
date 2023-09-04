package jsoner

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
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

func Quote(s any) string {
	return fmt.Sprintf("\"%s\"", s)
}

func TypeOf(obj any) reflect.Kind {
	return reflect.TypeOf(obj).Kind()
}

func processValue(value any) any {

	if TypeOf(value) == reflect.String {
		return Quote(value)
	} else {
		return value
	}
}

func (p person) insert() string {
	labels := []string{}
	values := []string{}

	// Use reflection to iterate over the fields of the struct
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// fmt.Printf("Field Name: %s, Field Value: %v\n", field.Name, processValue(value))
		fmt.Printf("Field Name: %s, Field Value: %v\n", field.Tag.Get("json"), processValue(value))

		labels = append(labels, fmt.Sprintf("%v", field.Tag.Get("json")))
		values = append(values, fmt.Sprintf("%v", processValue(value)))
	}

	insert_string := fmt.Sprintf("insert into person (%v) values (%v)", strings.Join(labels, ", "), strings.Join(values, ", "))

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
