package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"
)

func main() {

	// Verzeichnisse finden die für eine Löschung in Frage kommen

	// Gesamtgröße der einzelnen Verzeichnisse ermitteln -> ergibt sich aus der Summe der Größe der Dateien,
	// die es direkt oder indirekt enthält
	// Verzeichnisse selbst haben keine Größe

	// Suchen Sie zunächst alle Verzeichnisse mit einer Gesamtgröße von höchstens 100000 und berechnen Sie dann die Summe ihrer Gesamtgrößen

	data, _ := os.ReadFile("input.txt")

	cd := ""
	fs := map[string]int{} // Nil

	// Trimspace entfernt überflüssigen whitespace
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		fmt.Println("Ausgabe: ", line)

		// Prüfung ob String mit dem Prefix beginnt
		if strings.HasPrefix(line, "$ cd") {

			// Join --> ignoriert leere Elemente, damit wird das / entfernt
			cd = path.Join(cd, strings.Fields(line)[2]) //$ // []string slice
			fmt.Println("Wert von cd:", cd)

		} else if unicode.IsDigit([]rune(line)[0]) {
			var size int
			var name string
			fmt.Sscanf(line, "%d %s", &size, &name)
			fs[path.Join(cd, name)] = size
			fmt.Println("FS:", fs)
		}
	}
	fmt.Println("------------------------------------------------------")
	fmt.Println()

	//------------------------------------------------------------------------------

	ds := map[string]int{}
	for f, s := range fs {
		fmt.Println("Range Fs Map:", f, s)
		for d := path.Dir(f); d != "/"; d = path.Dir(d) {
			fmt.Println("DS[d]", ds[d])
			ds[d] += s
		}
		ds["/"] += s
		fmt.Println("DS/:", ds["/"])
	}

	//------------------------------------------------------------------------------

	var part1 int
	var part2 = ds["/"]
	// Iterieren durch alle Einträge in meiner Map
	for _, s := range ds {
		// Elemente durchgehen, die kleiner oder gleich 100000 sind.
		// Part 1
		if s <= 100000 {
			part1 += s
			fmt.Println("Ergebnis -->", part1)
		}
		// Part 2
		if s+70000000-ds["/"] >= 30000000 && s < part2 {
			part2 = s
		}
	}
	// Ausgabe : Part 1
	fmt.Println("Part 1:", part1)
	// Ausgabe : Part 2
	fmt.Println("Part 2:", part2)

}
