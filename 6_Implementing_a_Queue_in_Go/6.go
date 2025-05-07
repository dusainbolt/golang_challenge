package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Items []Flight
}

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (q *Queue) Pop() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue is Empty")
	}
	firstElemIndex := 0
	var flight Flight
	flight, q.Items = q.Items[firstElemIndex], q.Items[1:]
	return flight, nil
}

func (q *Queue) Push(flight Flight) {
	q.Items = append(q.Items, flight)
}

func (q *Queue) Peek() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue is Empty")
	}
	return q.Items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func main() {
	fmt.Println("Go Queue Implementation")
}
