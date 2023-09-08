package main

import (
	"log"
)

func main() {
	input := "abab"
	log.Println(longestPalindrome(input))
}

func longestPalindrome(s string) string {

	var low, maxLen int
	for i := 0; i < len(s); i++ {

		low1, maxLen1 := helper(s, i, i)
		low2, maxLen2 := helper(s, i, i+1)

		if maxLen2 > maxLen1 && maxLen2 > maxLen {
			maxLen = maxLen2
			low = low2
		} else if maxLen1 > maxLen {
			maxLen = maxLen1
			low = low1
		}
	}

	return s[low : low+maxLen+1]
}

func helper(s string, left, right int) (int, int) {

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	left++
	right--

	return left, right - left
}
