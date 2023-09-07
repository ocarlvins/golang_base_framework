package structs

import (
	"encoding/json"
	"webserver/utils"
)

type DbModels interface {
	ToJson() string
}

type Person struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
}

func (p Person) ToJson() string {
	u, err := json.Marshal(p)
	utils.Check(err)
	return string(u)
}
