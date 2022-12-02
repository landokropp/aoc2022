package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
)

func main() {
	err := realMain()
	if err != nil {
		panic(err)
	}
}

func realMain() error {

	//A Rock 1
	//B Paper 2
	//C Scissor 3

	//X Rock
	//Y Paper
	//Z Scissors

	//lose = 0
	//draw = 3
	// win = 6

	//Part 1
	// Win, Lose or Draw points / Cases
	score := make(map[string]int)
	score["A X"] = 3
	score["A Y"] = 6
	score["A Z"] = 0
	score["B X"] = 0
	score["B Y"] = 3
	score["B Z"] = 6
	score["C X"] = 6
	score["C Y"] = 0
	score["C Z"] = 3

	//Part2
	part2 := make(map[string]string)
	part2["A X"] = "A Z"
	part2["A Y"] = "A X"
	part2["A Z"] = "A Y"
	part2["B X"] = "B X"
	part2["B Y"] = "B Y"
	part2["B Z"] = "B Z"
	part2["C X"] = "C Y"
	part2["C Y"] = "C Z"
	part2["C Z"] = "C X"

	totalScore1 := 0
	totalScore2 := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		//Output wie in der Datei bis hier hin
		//-----------------------------------------------------------
		totalScore1 += score[line] + playerChoice(line)
		totalScore2 += score[part2[line]] + playerChoice(part2[line])
	}
	fmt.Println("Part 1:", totalScore1)
	fmt.Println("Part 2:", totalScore2)
	return err
}

func playerChoice(l string) int {
	switch l[len(l)-1:] {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		panic("Fehler!")
	}
}
