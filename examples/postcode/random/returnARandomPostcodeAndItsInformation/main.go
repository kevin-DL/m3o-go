package main

import (
	"fmt"
	"os"

	"go.m3o.com/postcode"
)

// Return a random postcode and its related info
func main() {
	postcodeService := postcode.NewPostcodeService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := postcodeService.Random(&postcode.RandomRequest{})
	fmt.Println(rsp, err)
}
