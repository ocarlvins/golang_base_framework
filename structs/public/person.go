package public

import (
	"encoding/json"
	"webserver/db"
	"webserver/utils"
)

type Person struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
}

func (p Person) TableArgs() db.TableArgs {
	return db.TableArgs{
		PrimaryKey:    []string{"id"},
		Schema:        "public",
		AutoIncrement: []string{"id"},
	}
}

func (p Person) ToJson() string {
	u, err := json.Marshal(p)
	utils.Check(err)
	return string(u)
}
