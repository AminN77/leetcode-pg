package main

func main() {

}

func characterReplacement(s string, k int) int {
	m := make(map[byte]int)
	var res int
	left := 0
	mostFre := 0
	for right := range s {
		m[s[right]]++
		if mostFre < m[s[right]] {
			mostFre = m[s[right]]
		}

		if (right-left+1)-mostFre > k {
			m[s[left]]--
			left += 1
		}

		if res < right-left+1 {
			res = right - left + 1
		}
	}

	return res
}
