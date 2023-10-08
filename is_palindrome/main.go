package main

import (
	"log"
	"strings"
)

func main() {
	input := "0P"
	log.Println(isPalindrome(input))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		for left < right && !isValid(s[left]) {
			left++
		}
		for left < right && !isValid(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

func isValid(c byte) bool { return c >= 'a' && c <= 'z' || c >= '0' && c <= '9' }
