package main

import (
	"fmt"
	"os"

	"go.m3o.com/sunnah"
)

// Get a list of available collections. A collection is
// a compilation of hadiths collected and written by an author.
func main() {
	sunnahService := sunnah.NewSunnahService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := sunnahService.Collections(&sunnah.CollectionsRequest{})
	fmt.Println(rsp, err)
}
