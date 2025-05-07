package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Flight struct which contains
// the origin, destination and price of a flight
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// IsSubset checks to see if the first set of
// flights is a subset of the second set of flights.
func IsSubset(first, second []Flight) bool {
	secondSet := make(map[string]bool)

	for _, flight := range second {
		hash := Hash(flight)
		secondSet[string(hash)] = true
	}

	for _, flight := range first {
		hash := Hash(flight)
		if !secondSet[string(hash)] {
			return false
		}
	}
	return true
}

func Hash(f Flight) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(f)
	return b.Bytes()
}

func main() {
	fmt.Println("Sets and Subsets Challenge")
	firstFlights := []Flight{
		{Origin: "GLA", Destination: "CDG", Price: 1000},
		{Origin: "GLA", Destination: "JFK", Price: 5000},
		{Origin: "GLA", Destination: "SNG", Price: 3000},
	}

	secondFlights := []Flight{
		{Origin: "GLA", Destination: "CDG", Price: 1000},
		{Origin: "GLA", Destination: "JFK", Price: 5000},
		{Origin: "GLA", Destination: "SNG", Price: 3000},
		{Origin: "GLA", Destination: "AMS", Price: 500},
	}

	subset := IsSubset(firstFlights, secondFlights)
	fmt.Println(subset)
}
