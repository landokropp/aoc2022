package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// TODO: Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
func main() {
	// öffnet den gewählten file name
	f, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// NewScanner returns a new Scanner to read from r. io.Reader
	scanner := bufio.NewScanner(f)

	// make built-in function allocates and initializes an object of type slice
	elves := make([]int, 1) // len und cap sind 1
	max := 0

	//Scan bringt den Scanner zum nächsten Token, die verfügt dann über die Methoden Bytes und Text
	for scanner.Scan() {

		// scanner.Text returns my text string
		line := scanner.Text()

		if len(line) > 0 {
			// integer Werte von meine txt
			calories, err := strconv.Atoi(line)
			// mehrere Schleifendurchgänge
			fmt.Println("Calories:", calories)
			if err != nil {
				panic(err)
			}

			// rechnet die Zahlen jedes Elfens zusammen
			elves[len(elves)-1] = elves[len(elves)-1] + calories
			fmt.Println("Ergebnis: ", elves[len(elves)-1])

			// max value von den Elfen
			if elves[len(elves)-1] > max {
				max = elves[len(elves)-1]
				fmt.Println("MAX:", max)
			}
		} else {

			elves = append(elves, 0)
			fmt.Println("elves: ", elves)
		}
	}
	fmt.Println("Teil 1:", max)

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	total := 0

	for _, element := range elves[:3] { // länge 3
		total = total + element
	}
	fmt.Println("elves3: ", elves[:3])
	fmt.Println("Teil 2:", total)
}
