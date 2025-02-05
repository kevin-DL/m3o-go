package main

import (
	"fmt"
	"os"

	"go.m3o.com/place"
)

// Search for places by text query
func main() {
	placeService := place.NewPlaceService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := placeService.Search(&place.SearchRequest{
		Location: "51.5074577,-0.1297515",
		Query:    "food",
	})
	fmt.Println(rsp, err)
}
