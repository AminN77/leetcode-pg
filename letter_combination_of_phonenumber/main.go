package main

import (
	"log"
)

var charMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

var gDigits string

func main() {
	input := "23"
	log.Println(letterCombinations(input))
}

func letterCombinations(digits string) []string {
	var res []string
	gDigits = digits
	if len(digits) > 0 {
		backTrack(0, "", &res)
	}

	return res
}

func backTrack(index int, curStr string, res *[]string) {
	if len(curStr) == len(gDigits) {
		*res = append(*res, curStr)
		return
	}
	for _, v := range charMap[gDigits[index]] {
		backTrack(index+1, curStr+v, res)
	}
}
