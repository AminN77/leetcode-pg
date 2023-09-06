package main

import (
	"log"
	"strings"
)

func main() {
	input := "abcabcabc"
	log.Println(lengthOfLongestSubstring_OP(input))
}

func lengthOfLongestSubstring(s string) int {
	length := 0
	dup := false
	if len(s) > 0 {
		length = 1
	}

	for i := 0; i < len(s); i++ {
		tempString := s[i : i+1]
		for j := i + 1; j < len(s); j++ {
			if strings.Contains(tempString, s[j:j+1]) {
				if len(tempString) > length {
					length = len(tempString)
					dup = true
				}
				break
			} else {
				tempString = s[i : j+1]
			}
		}

		if !dup {
			if len(tempString) > length {
				length = len(tempString)
			}
		}

		dup = false
	}

	return length
}

func lengthOfLongestSubstring_OP(s string) int {
	// Length of the given input string
	n := len(s)

	// Length of the longest substring
	var result int

	// Map to store visited characters along with their index
	charIndexMap := make(map[uint8]int)

	// start indicates the start of current substring
	var start int

	// Iterate through the string and slide the window each time you find a duplicate
	// end indicates the end of current substring
	for end := 0; end < n; end++ {

		// Check if the current character is a duplicate
		duplicateIndex, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			// Update the result for the substring in the current window before we found duplicate character
			result = max(result, end-start)

			// Remove all characters before the new
			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}

			// Slide the window since we have found a duplicate character
			start = duplicateIndex + 1
		}
		// Add the current character to hashmap
		charIndexMap[s[end]] = end
	}
	// Update the result for last window
	// For an input like "abc" the execution will not enter the above if statement for checking duplicates
	result = max(result, n-start)

	return result
}
