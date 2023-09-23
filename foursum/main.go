package main

import (
	"log"
	"sort"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	target := -1
	log.Println(fourSum(input, target))
}

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// three sum
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// two sum
			leftPointer := j + 1
			rightPointer := len(nums) - 1
			for leftPointer < rightPointer {
				sum := nums[i] + nums[j] + nums[leftPointer] + nums[rightPointer]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[leftPointer], nums[rightPointer]})
					leftPointer++
					for leftPointer < rightPointer && nums[leftPointer] == nums[leftPointer-1] {
						leftPointer++
					}
				} else if sum > target {
					rightPointer--
				} else {
					leftPointer++
				}
			}
		}
	}

	return res
}
