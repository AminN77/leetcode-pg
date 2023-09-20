package main

import (
	"log"
)

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.

func main() {
	input := 58
	log.Println(intToRoman(input))
}

func intToRoman(num int) string {
	return thousands(num) + hundreds(num) + tens(num) + ones(num)
}

func thousands(num int) string {
	count := num / 1000
	switch count {
	case 0:
		return ""
	case 1:
		return "M"
	case 2:
		return "MM"
	case 3:
		return "MMM"
	}

	return ""
}

func hundreds(num int) string {
	num = num % 1000
	count := num / 100

	switch count {
	case 0:
		return ""
	case 1:
		return "C"
	case 2:
		return "CC"
	case 3:
		return "CCC"
	case 4:
		return "CD"
	case 5:
		return "D"
	case 6:
		return "DC"
	case 7:
		return "DCC"
	case 8:
		return "DCCC"
	case 9:
		return "CM"
	}

	return ""
}

func tens(num int) string {
	num = num % 100
	count := num / 10

	switch count {
	case 0:
		return ""
	case 1:
		return "X"
	case 2:
		return "XX"
	case 3:
		return "XXX"
	case 4:
		return "XL"
	case 5:
		return "L"
	case 6:
		return "LX"
	case 7:
		return "LXX"
	case 8:
		return "LXXX"
	case 9:
		return "XC"
	}

	return ""
}

func ones(num int) string {
	count := num % 10
	switch count {
	case 0:
		return ""
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	}

	return ""
}
