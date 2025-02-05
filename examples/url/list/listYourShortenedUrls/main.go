package main

import (
	"fmt"
	"os"

	"go.m3o.com/url"
)

// List all the shortened URLs
func main() {
	urlService := url.NewUrlService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := urlService.List(&url.ListRequest{})
	fmt.Println(rsp, err)
}
