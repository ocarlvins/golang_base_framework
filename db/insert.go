package db

import (
	"fmt"
	"reflect"
	"strings"
	"webserver/utils"
)

func Insert(p any) string {
	var labels []string
	var values []string

	// Use reflection to iterate over the fields of the struct
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// fmt.Printf("Field Name: %s, Field Value: %v\n", field.Name, processValue(value))
		// fmt.Printf("Field Name: %s, Field Value: %v\n", field.Tag.Get("json"), utilities.QuoteIfString(value))

		if utils.NotNull(value) {
			labels = append(labels, field.Tag.Get("json"))
			values = append(values, utils.QuoteIfString(value))
		}
	}

	insert_string := fmt.Sprintf("insert into person (%v) values (%v)", strings.Join(labels, ", "), strings.Join(values, ", "))
	runQuery(insert_string)
	return insert_string
}
