package main

import (
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Create a record in the database. Optionally include an "id" field otherwise it's set automatically.
func main() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Create(&db.CreateRequest{
		Record: map[string]interface{}{
			"id":       "1",
			"name":     "Jane",
			"age":      42,
			"isActive": true,
		},
		Table: "example",
	})
	fmt.Println(rsp, err)
}
