package main

import (
	"log"
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	input := -123
	log.Println(reverse(input))
}

func reverse(x int) int {
	//check sign
	sign := true
	if x < 0 {
		sign = false
		x = x * -1
	}

	//reverse
	rev, _ := strconv.Atoi(revString(strconv.Itoa(x)))

	//check overflow
	if rev > math.MaxInt32 {
		return 0
	}

	//set sign
	if !sign {
		rev = rev * -1
	}

	return rev
}

func revString(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
