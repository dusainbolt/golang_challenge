package main

import "fmt"

// MinRotations tìm số lần mảng đã bị xoay
func MinRotations(array []int) int {
	n := len(array)
	low := 0
	high := n - 1

	for low <= high {
		// Nếu mảng đã được sắp xếp (không xoay), trả về low
		if array[low] <= array[high] {
			return low
		}

		mid := (low + high) / 2
		next := (mid + 1) % n
		prev := (mid - 1 + n) % n

		// Kiểm tra nếu phần tử mid là nhỏ nhất
		if array[mid] <= array[next] && array[mid] <= array[prev] {
			return mid
		}

		// Nếu array[mid] <= array[high], nghĩa là dãy phải đã được sắp xếp, tìm bên trái
		if array[mid] <= array[high] {
			high = mid - 1
		} else { // Ngược lại, tìm bên phải
			low = mid + 1
		}
	}

	return 0
}

func main() {
	fmt.Println("Min Rotation Challenge")

	testArr := []int{15, 18, 2, 3, 6, 12}
	min := MinRotations(testArr) // return 2
	fmt.Println(min)

	testArr2 := []int{7, 9, 11, 12, 5}
	min2 := MinRotations(testArr2) // return 4
	fmt.Println(min2)
}
