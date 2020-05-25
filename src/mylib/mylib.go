package mylib

import "fmt"

func add(l, r int) int {
	fmt.Printf("Adding %v and %v", l, r)
	return l + r
}

func sub(l, r int) int {
	return l - r
}
