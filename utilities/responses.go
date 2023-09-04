package utilities

import (
	"encoding/json"
	"net/http"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func JsonResponse(w http.ResponseWriter, Object any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	u, err := json.Marshal(Object)
	Check(err)
	w.Write(u)
}
