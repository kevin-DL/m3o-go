package main

import (
	"fmt"
	"os"

	"go.m3o.com/sunnah"
)

// Hadiths returns a list of hadiths and their corresponding text for a
// given book within a collection.
func main() {
	sunnahService := sunnah.NewSunnahService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := sunnahService.Hadiths(&sunnah.HadithsRequest{
		Book:       1,
		Collection: "bukhari",
	})
	fmt.Println(rsp, err)
}
