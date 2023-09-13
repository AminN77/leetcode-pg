package main

import (
	"log"
)

func main() {

	s := "abbbbab"
	p := "ab*ab"
	log.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	dp := map[int]bool{}
	l1, l2 := len(s), len(p)

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		key := 100*i + j
		if _, ok := dp[key]; ok {
			return dp[key]
		}
		if i >= l1 && j >= l2 {
			return true
		}
		if j >= l2 {
			return false
		}
		match := i < l1 && (s[i] == p[j] || p[j] == '.')
		if j+1 < l2 && p[j+1] == '*' {
			dp[key] = dfs(i, j+2) || (match && dfs(i+1, j))
			return dp[key]
		}
		if match {
			dp[key] = dfs(i+1, j+1)
			return dp[key]
		}
		dp[key] = false
		return false
	}
	return dfs(0, 0)
}
