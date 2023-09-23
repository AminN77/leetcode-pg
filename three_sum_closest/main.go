package main

import (
	"log"
	"math"
	"sort"
)

func main() {
	input := []int{-1, 2, 1, -4}
	target := 1
	log.Println(threeSumClosest(input, target))
}

func threeSumClosest(nums []int, target int) int {
	diff := math.MaxInt32
	sum := 0
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		leftPointer := i + 1
		rightPointer := len(nums) - 1
		for leftPointer < rightPointer {
			tempSum := nums[i] + nums[leftPointer] + nums[rightPointer]
			if tempSum == target {
				return target
			} else if tempSum > target {
				rightPointer--
			} else {
				leftPointer++
			}

			if int(math.Abs(float64(tempSum-target))) < diff {
				diff = int(math.Abs(float64(tempSum - target)))
				sum = tempSum
			}
		}
	}

	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		for k := j + 1; k < len(nums); k++ {
	//			tempDiff := nums[i] + nums[j] + nums[k] - target
	//			if int(math.Abs(float64(tempDiff))) < diff {
	//				diff = int(math.Abs(float64(tempDiff)))
	//				sum = nums[i] + nums[j] + nums[k]
	//			}
	//		}
	//	}
	//}

	return sum
}
