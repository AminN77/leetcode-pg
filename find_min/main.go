package main

import "log"

func main() {
	input := []int{3, 4, 5, 6, 7, 1, 2}
	log.Println(findMin(input))
}

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if mid > 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nums[0]
}
