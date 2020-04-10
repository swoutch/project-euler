package main

import (
	"strconv"
	"math"
	"fmt"
)

func IsPalindrome(n int) bool {
	result := true
	digits := strconv.Itoa(n)
	i := 0
	j := len(digits) - 1
	for i < j && result {
		if digits[i] != digits[j] {
			result = false
		}
		i++
		j--
	}
	return result
}

func FindLargestPalindrome(numberOfDigits int) int {
	var largestPalindrome int
	for factor1:=1; factor1<int(math.Pow(10, float64(numberOfDigits))); factor1++ {
		for factor2:=1; factor2<int(math.Pow(10, float64(numberOfDigits))); factor2++ {
			product := factor1*factor2
			if IsPalindrome(product) && largestPalindrome < product {
				largestPalindrome = product
			}
		}
	}
	return largestPalindrome
}

func main() {
	fmt.Println(FindLargestPalindrome(3))
}
