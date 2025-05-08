package main

import "fmt"

type MyInt int

func (n *MyInt) IsArmstrong() bool {
	num := int(*n)
	sum := 0
	temp := num

	// 1. Đếm số chữ số
	count := 0
	for t := num; t > 0; t /= 10 {
		count++
	}

	// 2. Lặp qua từng chữ số, tính lũy thừa và cộng vào tổng
	for temp > 0 {
		digit := temp % 10 // Lấy chữ số cuối
		power := 1
		for i := 0; i < count; i++ {
			power *= digit // Tính digit^count
		}
		sum += power
		temp /= 10 // Bỏ chữ số cuối
	}

	//  So sánh tổng với số ban đầu
	return sum == num
}

func main() {
	fmt.Println("Armstrong Numbers")

	var num1 MyInt = 371
	fmt.Println(num1.IsArmstrong())
}
