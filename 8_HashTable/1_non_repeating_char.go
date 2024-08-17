package main

import (
	"fmt"
)

type Solution struct{}

func (s *Solution) firstUniqChar(sStr string) int {

	// Create a Hashmap
	freq := make(map[rune]int)

	// Populate the hashmap
	for _, s := range sStr {
		freq[s]++
	}

	for idx, c := range sStr {
		if freq[c] == 1 {
			return idx
		}
	}

	return -1
}

func main() {
	sol := &Solution{}

	fmt.Println(sol.firstUniqChar("apple")) // Expected: 0
	fmt.Println(sol.firstUniqChar("abcab")) // Expected: 2
	fmt.Println(sol.firstUniqChar("abab"))  // Expected: -1
}
