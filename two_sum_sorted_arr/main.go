package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			break
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{left + 1, right + 1}
}
