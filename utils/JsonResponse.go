package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, Object any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	u, err := json.Marshal(Object)
	Check(err)
	_, err = w.Write(u)
	if err != nil {
		panic(err)
	}
}
