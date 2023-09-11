package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := ""
	log.Println(myAtoi(input))
}

func myAtoi(s string) int {
	if s == "" || strings.Trim(s, " ") == "" {
		return 0
	}
	// 1. ignore spaces
	s = strings.Trim(s, " ")

	// 2. check sign
	sign := true
	if !unicode.IsDigit(rune(s[0])) {
		if s[0:1] == "-" {
			sign = false
		} else if s[0:1] == "+" {
		} else {
			return 0
		}

		s = s[1:]
	}

	// 3. parse digits
	count := 0
	for i := 0; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			break
		} else {
			count++
		}
	}

	// 4. extract digits and convert to int
	s = s[0:count]
	out, _ := strconv.Atoi(s)

	// 5. assign sign
	if !sign {
		out = out * -1
	}

	// 6. check boundaries
	if out > math.MaxInt32 {
		return math.MaxInt32
	}

	if out < math.MinInt32 {
		return math.MinInt32
	}

	return out
}
