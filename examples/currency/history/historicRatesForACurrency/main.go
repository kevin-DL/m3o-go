package main

import (
	"fmt"
	"os"

	"go.m3o.com/currency"
)

// Returns the historic rates for a currency on a given date
func main() {
	currencyService := currency.NewCurrencyService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := currencyService.History(&currency.HistoryRequest{
		Code: "USD",
		Date: "2021-05-30",
	})
	fmt.Println(rsp, err)
}
