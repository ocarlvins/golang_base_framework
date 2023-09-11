package utils

import (
	"fmt"
	"reflect"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Quote(s any) string {
	return fmt.Sprintf("'%s'", s)
}

func TypeOf(obj any) reflect.Kind {
	return reflect.TypeOf(obj).Kind()
}

func QuoteIfString(value any) string {
	if TypeOf(value) == reflect.String {
		return Quote(value)
	} else {
		return fmt.Sprintf("%v", value)
	}
}

func FieldType(value any) string {
	switch TypeOf(value) {
	case reflect.String:
		return "varchar"
	case reflect.Int:
		return "integer"
	case reflect.Float32:
		return "numeric"
	case reflect.Bool:
		return "boolean"
	default:
		return "varchar"
	}
}

func IsNull(val any) bool {
	if TypeOf(val) == reflect.String && val == "" {
		return true
	}
	return val == nil
}

func NotNull(val any) bool {
	return !IsNull(val)
}

func StringContains(stringsArray []string, targetString string) bool {
	found := false
	for _, str := range stringsArray {
		if str == targetString {
			found = true
			break
		}
	}

	if found {
		return true
	} else {
		return false
	}
}
