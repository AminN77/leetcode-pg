package main

import "math"

func main() {

}

func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}
	tMap := make(map[byte]int)
	for i := range t {
		tMap[t[i]] = tMap[t[i]] + 1
	}
	window := make(map[byte]int)

	have, need := 0, len(tMap)
	res := []int{-1, -1}
	resLen := math.MaxUint32
	l := 0
	for r := range s {
		c := s[r]
		window[c] = window[c] + 1
		if window[c] == tMap[c] {
			have++
			//just satisfied a condition
		}
		for have == need {
			//requirement met, see if it's a new min
			if r-l+1 < resLen {
				res[0] = l
				res[1] = r
				resLen = r - l + 1
			}
			//shrink window
			window[s[l]] = window[s[l]] - 1
			if window[s[l]] < tMap[s[l]] {
				have--
			}
			l++
		}
	}
	if resLen < math.MaxUint32 {
		l = res[0]
		r := res[1]
		return s[l : r+1]
	} else {
		return ""
	}
	return ""
}
