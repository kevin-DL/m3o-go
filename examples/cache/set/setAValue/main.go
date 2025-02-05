package main

import (
	"fmt"
	"os"

	"go.m3o.com/cache"
)

// Set an item in the cache. Overwrites any existing value already set.
func main() {
	cacheService := cache.NewCacheService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := cacheService.Set(&cache.SetRequest{
		Key:   "foo",
		Value: "bar",
	})
	fmt.Println(rsp, err)
}
