package main

import "fmt"

func main() {
	r := []rune("Hello,\tworld!")
	p := &r[7]
	*p = 'G'
	fmt.Print(r[7:8])
}
