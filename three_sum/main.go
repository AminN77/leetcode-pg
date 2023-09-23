package main

import (
	"log"
	"sort"
)

// [-1 0 1]

func main() {
	input := []int{2, 2, 2, 2}
	log.Println(threeSum(input))
}

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// two sum
		leftPointer := i + 1
		rightPointer := len(nums) - 1
		for leftPointer < rightPointer {
			sum := nums[i] + nums[leftPointer] + nums[rightPointer]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[leftPointer], nums[rightPointer]})
				leftPointer++
				for leftPointer < rightPointer && nums[leftPointer] == nums[leftPointer-1] {
					leftPointer++
				}
			} else if sum > 0 {
				rightPointer--
			} else {
				leftPointer++
			}
		}
	}

	return res
}
