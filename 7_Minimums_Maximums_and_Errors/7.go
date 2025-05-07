package main

import (
	"errors"
	"fmt"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func GetMinMax(flights []Flight) (int, int, error) {
	// Implement me :)
	if len(flights) == 0 {
		return 0, 0, errors.New("no flights provided")

	}

	min := flights[0].Price
	max := flights[0].Price

	for _, flight := range flights {
		if flight.Price < min {
			min = flight.Price
		}
		if flight.Price > max {
			max = flight.Price
		}
	}

	return min, max, nil
}

func main() {
	fmt.Println("Getting the Minimum and Maximum Flight Prices")

	flights := []Flight{
		{Price: 100},
		{Price: 50},
		{Price: 70},
		{Price: 20},
		{Price: 90},
		{Price: 89},
	}

	min, max, err := GetMinMax(flights)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Min: %d, Max: %d\n", min, max)
	}
}
