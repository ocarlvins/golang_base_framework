package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Quote(s any) string {
	return fmt.Sprintf("\"%s\"", s)
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

func JsonResponse(w http.ResponseWriter, Object any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	u, err := json.Marshal(Object)
	Check(err)
	w.Write(u)
}
