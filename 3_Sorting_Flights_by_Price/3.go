package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type ByPrice []Flight

func (p ByPrice) Len() int {
	return len(p)
}

func (p ByPrice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByPrice) Less(i, j int) bool {
	return p[i].Price > p[j].Price
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	sort.Sort(ByPrice(flights))
	return flights

}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	// an empty slice of flights
	flights := []Flight{
		{Price: 100},
		{Price: 50},
		{Price: 70},
		{Price: 20},
		{Price: 90},
		{Price: 89},
	}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
