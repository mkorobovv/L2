package main

import (
	"fmt"
	"sort"
)

func isAnagram(s1 []rune, s2 []rune) bool {
	if len(s1) != len(s2) {
		return false
	}
	if sort.Sort(s1) == sort.Sort(s2) {
		return true
	}
	return false
}

func main() {
	fmt.Println()
}
