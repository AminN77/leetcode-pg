package main

import (
	"log"
	"math"
)

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	log.Println(maxArea2(input))
}

// maxArea from O(n^2) brute force solution
func maxArea(height []int) int {
	maximum := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			water := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
			if water > maximum {
				maximum = water
			}
		}
	}

	return maximum
}

// maxArea2 from O(n) two pointers approach
func maxArea2(height []int) int {
	maximum := 0
	leftPointer := 0
	rightPointer := len(height) - 1
	for leftPointer != rightPointer {
		water := (rightPointer - leftPointer) * localMin(height[leftPointer], height[rightPointer])
		if water > maximum {
			maximum = water
		}

		if height[leftPointer] > height[rightPointer] {
			rightPointer--
		} else {
			leftPointer++
		}
	}

	return maximum
}

func localMin(i int, j int) int {
	if i > j {
		return i
	}

	return j
}
