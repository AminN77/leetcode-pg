package main

import "log"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	log.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums3 []int
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i++
		} else {
			nums3 = append(nums3, nums2[j])
			j++
		}
	}

	// Append remaining elements from nums1 (if any)
	for i < len(nums1) {
		nums3 = append(nums3, nums1[i])
		i++
	}

	// Append remaining elements from nums2 (if any)
	for j < len(nums2) {
		nums3 = append(nums3, nums2[j])
		j++
	}

	length := len(nums3)

	if length%2 == 1 {
		return float64(nums3[length/2])
	}

	mid1 := nums3[length/2-1]
	mid2 := nums3[length/2]
	return float64(mid1+mid2) / 2.0
}
