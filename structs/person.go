package structs

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"webserver/utilities"
)

type Person struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
}

func (p Person) ToJson() string {
	u, err := json.Marshal(p)
	utilities.Check(err)
	return string(u)
}

func (p Person) Insert() string {
	labels := []string{}
	values := []string{}

	// Use reflection to iterate over the fields of the struct
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// fmt.Printf("Field Name: %s, Field Value: %v\n", field.Name, processValue(value))
		fmt.Printf("Field Name: %s, Field Value: %v\n", field.Tag.Get("json"), utilities.QuoteIfString(value))

		labels = append(labels, field.Tag.Get("json"))
		values = append(values, utilities.QuoteIfString(value))
	}

	insert_string := fmt.Sprintf("insert into person (%v) values (%v)", strings.Join(labels, ", "), strings.Join(values, ", "))

	return insert_string
}
