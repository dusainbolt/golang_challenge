package main

import "fmt"

// OddEvenFactors kiểm tra xem một số có số lượng ước là chẵn hay lẻ.
// Nếu là số chính phương (square number), số lượng ước sẽ lẻ.
// Ngược lại, các số còn lại sẽ có số ước chẵn (do ước luôn đi theo cặp: i và num/i).
func OddEvenFactors(num int) string {
	count := 0

	// Duyệt từ 1 đến căn bậc hai của num
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			// Nếu i * i == num thì chỉ có 1 ước (vì i và num/i trùng nhau)
			if i*i == num {
				count++ // chỉ cộng 1
			} else {
				count += 2 // có cả i và num/i là 2 ước khác nhau
			}
		}
	}

	fmt.Println(count)
	// Nếu tổng số ước là chẵn → "even", ngược lại → "odd"
	if count%2 == 0 {
		return "even"
	}
	return "odd"
}

func main() {
	fmt.Println("Odd or Even Factors")

	numFactors := OddEvenFactors(23)
	fmt.Println(numFactors) // "even"

	numFactors = OddEvenFactors(36)
	fmt.Println(numFactors) // "odd"
}
