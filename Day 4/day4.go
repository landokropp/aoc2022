package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	input, _ := os.ReadFile("input.txt")
	part1, part2 := 0, 0

	for _, s := range strings.Fields(strings.TrimSpace(string(input))) {
		var s1, e1, s2, e2 int
		fmt.Sscanf(s, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)
		//fmt.Println("Input:",s)

		if s1 <= s2 && e1 >= e2 || s1 >= s2 && e1 <= e2 {
			part1++
		}
		if s1 <= e2 && e1 >= s2 {
			part2++
		}
	}

	fmt.Println("Ergebnis Part1:", part1)
	fmt.Println("Ergebnis Part2:", part2)
}
