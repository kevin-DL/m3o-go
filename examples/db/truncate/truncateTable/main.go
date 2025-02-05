package main

import (
	"fmt"
	"os"

	"go.m3o.com/db"
)

// Truncate the records in a table
func main() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := dbService.Truncate(&db.TruncateRequest{
		Table: "example",
	})
	fmt.Println(rsp, err)
}
