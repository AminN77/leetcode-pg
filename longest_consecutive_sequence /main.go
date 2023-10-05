package main

import (
	"log"
	"sort"
)

func main() {
	input := []int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7}
	log.Println(longestConsecutive(input))
}

func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	maximum := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			if nums[i] == nums[i-1] {
				continue
			}
			if nums[i]-nums[i-1] == 1 {
				temp += 1
			} else {
				if maximum < temp {
					maximum = temp
				}
				temp = 1
			}
		} else {
			temp += 1
		}
	}

	if maximum < temp {
		return temp
	}

	return maximum
}
