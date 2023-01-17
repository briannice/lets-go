package main

import (
	"fmt"
	"sync"
)

// Package level WaitGroup.
var wg sync.WaitGroup

// Map of countries and their capital names.
var countriesWithCapitals = map[string]string{
	"Germany":     "Berlin",
	"France":      "Paris",
	"Italy":       "Rome",
	"Belgium":     "Brussels",
	"Netherlands": "Amsterdam",
	"Denmark":     "Kopenhagen",
	"Norway":      "Oslo",
	"Sweden":      "Stockholm",
	"Russia":      "Moscow",
	"Ukraine":     "Kiev",
	"Switserland": "Zurich",
	"Austria":     "Vienna",
}

func searchCapitalByCountry(i int, country string) {

	// Decrements the WaitGroup counter by one when the function returns.
	defer wg.Done()

	// Search the capital of the given country. Also check if the country
	// exists in the list of countries.
	capital, found := countriesWithCapitals[country]
	if found {
		fmt.Printf("%d) The capital of %s is %s\n", i, country, capital)
	} else {
		fmt.Printf("%d) The capital of %s is not found\n", i, country)
	}
}

func main() {

	// List of countries of which we want to find the capital.
	countries := []string{
		"Germany",
		"Russia",
		"Spain",
		"Netherlands",
		"Belgium",
	}

	// Search the capital of each country using a Go routine. Add 1 to the
	// WaitGroup counter before starting the Go routing to ensure the program
	// will not exit until all Go routines are done.
	for i, country := range countries {
		wg.Add(1)
		go searchCapitalByCountry(i, country)
	}

	// Wait for all Go routines to finish.
	wg.Wait()
}
