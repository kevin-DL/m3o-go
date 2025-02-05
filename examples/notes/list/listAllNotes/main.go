package main

import (
	"fmt"
	"os"

	"go.m3o.com/notes"
)

// List all the notes
func main() {
	notesService := notes.NewNotesService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := notesService.List(&notes.ListRequest{})
	fmt.Println(rsp, err)
}
