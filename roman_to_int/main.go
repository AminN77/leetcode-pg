package main

import "log"

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

var alpMap = map[uint8]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func main() {
	input := "LVIII"
	log.Println(romanToInt(input))
}

func romanToInt(s string) int {
	length := len(s)

	if length == 0 {
		return 0
	}

	if length == 1 {
		return alpMap[s[0]]
	}

	sum := alpMap[s[length-1]]
	for i := length - 2; i >= 0; i-- {
		if alpMap[s[i]] < alpMap[s[i+1]] {
			sum -= alpMap[s[i]]
		} else {
			sum += alpMap[s[i]]
		}
	}

	return sum
}
