package main

import (
	"fmt"
	"os"
)

func uniqueCharacters(s string) bool {

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i:i+1] == s[j:j+1] {
				return false
			}
		}
	}

	return true
}

func main() {
	// Format input
	dat, _ := os.ReadFile("input.txt")
	input := string(dat)

	//Part One
	marker := input[:4]

	for i := len(marker); i < len(input); i++ {
		if uniqueCharacters(marker) {
			fmt.Println("Part one Answer: ", i)
			break
		} else {
			marker = marker[1:] + input[i:i+1]
		}
	}
	


	// part Two
	markerb := input[:14]

	for i := len(markerb); i < len(input); i++ {
		if uniqueCharacters(markerb) {
			fmt.Println("Part Two Answer: ", i)
			break
		} else {
			markerb = markerb[1:] + input[i:i+1]
		}
	}
}