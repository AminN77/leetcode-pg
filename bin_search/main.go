package main

import "log"

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	target := 9
	log.Println(search(input, target))
}

func search(nums []int, target int) int {
	for left, right := 0, len(nums)-1; left <= right; {
		pos := left + (right-left)/2
		val := nums[pos]
		if val == target {
			return pos
		} else if val < target {
			left = pos + 1
		} else /* if val > target */ {
			right = pos - 1
		}
	}
	return -1
}
