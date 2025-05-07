package main

import (
	"errors"
	"fmt"
)

// Flight struct represents a flight with origin, destination, and price.
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// Stack struct represents a stack data structure for Flights.
type Stack struct {
	Items []Flight
}

// Push adds a Flight to the top of the stack.
func (s *Stack) Push(f Flight) {
	s.Items = append(s.Items, f)
}

// Pop removes and returns the top Flight from the stack.
// Returns an error if the stack is empty.
func (s *Stack) Pop() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("stack is empty")
	}
	topIndex := len(s.Items) - 1
	top := s.Items[topIndex]
	s.Items = s.Items[:topIndex]
	return top, nil
}

func (s *Stack) Peek() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("stack is empty")
	}
	return s.Items[len(s.Items)-1], nil
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func main() {
	fmt.Println("Go Stack Implementation")

	stack := &Stack{}

	// Push some flights onto the stack
	stack.Push(Flight{Price: 500})
	stack.Push(Flight{Price: 700})
	stack.Push(Flight{Price: 600})

	// Peek at the top flight
	if topFlight, err := stack.Peek(); err == nil {
		fmt.Printf("Top Flight: %+v\n", topFlight)
	} else {
		fmt.Println("Error:", err)
	}

	// Pop all flights from the stack
	for !stack.IsEmpty() {
		if flight, err := stack.Pop(); err == nil {
			fmt.Printf("Popped Flight: %+v\n", flight)
		} else {
			fmt.Println("Error:", err)
		}
	}

	// Attempt to pop from an empty stack
	if _, err := stack.Pop(); err != nil {
		fmt.Println("Error:", err)
	}
}
