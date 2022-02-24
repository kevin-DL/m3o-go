package main

import (
	"fmt"
	"os"

	"go.m3o.com/contact"
)

// Read contact details
func main() {
	contactService := contact.NewContactService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := contactService.Read(&contact.ReadRequest{
		Id: "42e48a3c-6221-11ec-96d2-acde48001122",
	})
	fmt.Println(rsp, err)
}
