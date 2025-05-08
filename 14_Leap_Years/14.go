package main

import "fmt"

func CheckLeapYear(year int) bool {
	// Năm nhuận là năm chia hết cho 4
	// nhưng nếu chia hết cho 100 thì phải chia hết cho 400
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
	fmt.Println("Check Leap Year")

	year := 2020
	fmt.Println(CheckLeapYear(year))
}
