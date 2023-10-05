package main

import (
	"log"
)

func main() {
	input := []int{1, 2, 3, 4}
	log.Println(productExceptSelf(input))
}

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	for i := range output {
		output[i] = 1
	}

	// applying prefix
	pref := 1
	for i := range nums {
		output[i] = output[i] * pref
		pref = pref * nums[i]
	}

	// applying postfix
	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] = output[i] * post
		post = post * nums[i]
	}

	return output
}
