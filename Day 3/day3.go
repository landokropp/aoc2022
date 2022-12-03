package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// prio map
// BONUS: map daten dynamisch generieren
// int -> char
// for (0-25 -> i )
// 97 + i(0) -> a
// 97 + i(1) -> b ...
// for(0-25 -> i)

var ergebnis int
var s1 string
var s2 string

func main() {

	lowcases := make(map[string]int)
	lowcases["a"] = 1
	lowcases["b"] = 2
	lowcases["c"] = 3
	lowcases["d"] = 4
	lowcases["e"] = 5
	lowcases["f"] = 6
	lowcases["g"] = 7
	lowcases["h"] = 8
	lowcases["i"] = 9
	lowcases["j"] = 10
	lowcases["k"] = 11
	lowcases["l"] = 12
	lowcases["m"] = 13
	lowcases["n"] = 14
	lowcases["o"] = 15
	lowcases["p"] = 16
	lowcases["q"] = 17
	lowcases["r"] = 18
	lowcases["s"] = 19
	lowcases["t"] = 20
	lowcases["u"] = 21
	lowcases["v"] = 22
	lowcases["w"] = 23
	lowcases["x"] = 24
	lowcases["y"] = 25
	lowcases["z"] = 26
	lowcases["A"] = 27
	lowcases["B"] = 28
	lowcases["C"] = 29
	lowcases["D"] = 30
	lowcases["E"] = 31
	lowcases["F"] = 32
	lowcases["G"] = 33
	lowcases["H"] = 34
	lowcases["I"] = 35
	lowcases["J"] = 36
	lowcases["K"] = 37
	lowcases["L"] = 38
	lowcases["M"] = 39
	lowcases["N"] = 40
	lowcases["O"] = 41
	lowcases["P"] = 42
	lowcases["Q"] = 43
	lowcases["R"] = 44
	lowcases["S"] = 45
	lowcases["T"] = 46
	lowcases["U"] = 47
	lowcases["V"] = 48
	lowcases["W"] = 49
	lowcases["X"] = 50
	lowcases["Y"] = 51
	lowcases["Z"] = 52

	// Input aufteilen in Zeilen + jede Zeile in 2 hÃ¤lften teilen
	for _, line := range strings.Split(input, "\n") {

		split1 := line[:len(line)/2]
		split2 := line[len(line)/2:]
		fmt.Println("HalfOne:", split1)
		fmt.Println("HalfTwo:", split2)
		fmt.Println()

		// jede hÃ¤lfte nach gleichen Buchstaben durchsuchen
		// -> z.B. ergebis: e

		var done = false

		for _, v1 := range split1 {
			s1 = string(v1)

			// Alle zeichen aus split2 durchgehen und vergleichen

			for _, v2 := range split2 {
				s2 = string(v2)

				// log
				//fmt.Println("Buchstabe:",s1, s2)

				// check ... s1 ==s2
				if s1 == s2 {
					fmt.Println("check...:", s1, s2)

					if done == false {

						// buchstabe nachschlagen welche Prio Zahl
						prio1 := lowcases[s1]

						// log
						fmt.Println("prio1...:", prio1)

						// ergbnis += prio Zahl
						ergebnis = ergebnis + prio1

						done = true
					}

				}

				// Part two...
			}

		}

	}
	// ausgabe: summe prio zahlen
	fmt.Println("Total: ", ergebnis)
}
