package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// Kiểm tra xem một số có phải số nguyên tố không
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Hoán vị slice số nguyên - sử dụng backtracking
func permute(nums []int, result *[]int) {
	var generate func(int)
	generate = func(n int) {
		if n == 1 {
			// chuyển slice số thành số nguyên, thêm vào danh sách kết quả
			val, _ := strconv.Atoi(joinInts(nums))
			*result = append(*result, val)
			return
		}
		for i := 0; i < n; i++ {
			generate(n - 1)
			if n%2 == 1 {
				nums[0], nums[n-1] = nums[n-1], nums[0]
			} else {
				nums[i], nums[n-1] = nums[n-1], nums[i]
			}
		}
	}
	generate(len(nums))
}

// Chuyển slice số nguyên thành chuỗi
func joinInts(nums []int) string {
	str := ""
	for _, v := range nums {
		str += strconv.Itoa(v)
	}
	return str
}

// Tìm số nguyên tố pandigital lớn nhất
func LargestPandigitalPrime() int {
	for n := 7; n >= 1; n-- { // chỉ cần xét từ 7 trở xuống
		digits := make([]int, n)
		for i := 0; i < n; i++ {
			digits[i] = i + 1
		}

		var permutations []int
		permute(digits, &permutations)

		// Sắp xếp giảm dần để tìm số lớn nhất trước
		sort.Slice(permutations, func(i, j int) bool {
			return permutations[i] > permutations[j]
		})

		for _, p := range permutations {
			if isPrime(p) {
				return p
			}
		}
	}

	return 0
}

func main() {
	fmt.Println("Pandigital Primes")

	pandigitalPrime := LargestPandigitalPrime()
	fmt.Println(pandigitalPrime) // output: 7652413
}
