package main

import (
	"log"
	"strings"
)

func main() {
	input := []string{"flower", "flow", "flight"}
	log.Println(longestCommonPrefix(input))
}

func longestCommonPrefix(strs []string) string {
	pref := strs[0]
	for i := range strs {
		for {
			if strings.HasPrefix(strs[i], pref) {
				break
			}
			pref = pref[0 : len(pref)-1]
		}
	}

	return pref
}
