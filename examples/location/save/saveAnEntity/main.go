package main

import (
	"fmt"
	"os"

	"go.m3o.com/location"
)

// Save an entity's current position
func main() {
	locationService := location.NewLocationService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := locationService.Save(&location.SaveRequest{
		Entity: &location.Entity{
			Id: "1",
			Location: &location.Point{
				Latitude:  51.511061,
				Longitude: -0.120022,
				Timestamp: 1622802761,
			},
			Type: "bike",
		},
	})
	fmt.Println(rsp, err)
}
