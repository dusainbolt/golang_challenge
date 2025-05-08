package main

import (
	"fmt"
	"math"
)

func DiffSquares(n, m int) int {
	n1 := math.Pow(float64(n), 2)
	n2 := math.Pow(float64(m), 2)
	return int(n1) - int(n2)
}

func main() {
	fmt.Println("Calculate The Difference of Squares")
	result := DiffSquares(5, 4)
	fmt.Println(result)
}
