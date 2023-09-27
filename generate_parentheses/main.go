package main

import (
	"log"
	"strings"
)

func main() {
	input := 3
	log.Println(generateParenthesis(input))
}

func generateParenthesis(n int) []string {
	var res []string
	var stack []string
	backTrack(0, 0, n, &res, &stack)
	return res
}

func backTrack(open, close, n int, res, stack *[]string) {
	if open == close && open == n {
		*res = append(*res, strings.Join(*stack, ""))
		return
	}

	if open < n {
		*stack = append(*stack, "(")
		backTrack(open+1, close, n, res, stack)
		*stack = (*stack)[:len(*stack)-1]
	}

	if close < open {
		*stack = append(*stack, ")")
		backTrack(open, close+1, n, res, stack)
		*stack = (*stack)[:len(*stack)-1]
	}

}
