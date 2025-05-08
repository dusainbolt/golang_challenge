package main

import "fmt"

// Số tam giác thứ n là tổng của n số nguyên dương đầu tiên:
// T(n) = 1 + 2 + 3 + ... + n
//
// Ví dụ:
// T(1) = 1
// T(2) = 1 + 2 = 3
// T(3) = 1 + 2 + 3 = 6
// T(4) = 1 + 2 + 3 + 4 = 10
//
// Ta có thể tính nhanh bằng công thức:
// T(n) = n * (n + 1) / 2
func TriangularNumber(n int) int {
	return (n * (n + 1)) / 2
}

func main() {
	fmt.Println("Returning the 'nth' triangular number")

	number := TriangularNumber(3)
	fmt.Println(number) // '6'

}
