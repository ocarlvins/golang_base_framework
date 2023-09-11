package db

import (
	"fmt"
	"reflect"
	"strings"
	"webserver/utils"
)

type TableArgs struct {
	PrimaryKey    []string
	Schema        string
	AutoIncrement []string
}

type DbModels interface {
	TableArgs() TableArgs
}

func CreateString(p DbModels) string {
	var fields []string

	// Use reflection to iterate over the fields of the struct
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	table_args := p.TableArgs()
	schema := table_args.Schema
	if schema == "" {
		schema = "public"
	}
	primary_keys := table_args.PrimaryKey
	autoincrement := table_args.AutoIncrement

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// fmt.Printf("Field Name: %s, Field Value: %v\n", field.Name, processValue(value))
		// fmt.Printf("Field Name: %s, Field Value: %v\n", field.Tag.Get("json"), utilities.QuoteIfString(value))

		label := field.Tag.Get("json")
		field_type := utils.FieldType(value)

		if utils.StringContains(primary_keys, label) {
			if utils.StringContains(autoincrement, label) {
				field_type = fmt.Sprintf(" serial primary key")
			} else {
				field_type = fmt.Sprintf(" %v primary key", field_type)
			}
		}
		fields = append(fields, fmt.Sprintf(" %v %v", label, field_type))
	}

	create_string := fmt.Sprintf("create table %v (\n%v\n)", reflect.TypeOf(p), strings.Join(fields, ",\n"))
	return create_string
}

func Create(p DbModels) string {
	create_string := CreateString(p)
	fmt.Println(create_string)
	runQuery(create_string)
	fmt.Println("Table created successfully")
	return "Table created successfully"
}
