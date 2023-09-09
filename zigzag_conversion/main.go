package main

import "log"

func main() {
	input := "paypalishiring"
	log.Println(convert(input, 3))
}

// convert a string to it's zigzag representation.
// Although two nested loops have been used, The time complexity is from O(N).
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := ""
	for r := 0; r < numRows; r++ {
		inc := 2 * (numRows - 1)
		for i := r; i < len(s); i += inc {
			// appending characters from same rows by jumping inc<2*rows-1> size steps
			res += s[i : i+1]

			// extra characters for inner rows
			if r > 0 && r < numRows-1 && i+inc-2*r < len(s) {
				res += s[i+inc-2*r : i+inc-2*r+1]
			}
		}
	}

	return res
}
