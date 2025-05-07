package main

import "fmt"

func CheckPermutations(str1, str2 string) bool {
	// implement
	if len(str1) != len(str2) {
		return false
	}

	charCountMap := make(map[rune]int)

	for _, char := range str1 {
		charCountMap[char]++
	}
	fmt.Println(charCountMap)

	for _, char := range str2 {
		if _, exists := charCountMap[char]; !exists {
			return false
		}
		charCountMap[char]--
		if charCountMap[char] < 0 {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println("Check Permutations Challenge")

	str1 := "adcme"
	str2 := "medac"

	isPermutation := CheckPermutations(str1, str2)
	fmt.Println(isPermutation)

}
