package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	total := 0
	fmt.Print("------------\n\n\n")

	line := ""
	prevLine := ""
	prevPrevLine := ""

	i := 0
	for fileScanner.Scan() {
		prevPrevLine = prevLine
		prevLine = line
		line = fileScanner.Text()
		if i%3 == 2 {
			fmt.Printf("%s\n", prevPrevLine)
			fmt.Printf("%s\n", prevLine)
			fmt.Printf("%s\n", line)
			fmt.Printf("%d\n", i)
			fmt.Printf("\b")
			// compute

			firstChars := make(map[rune]bool)

			for _, c := range prevPrevLine {
				firstChars[c] = true
			}

			potentialCommon := make(map[rune]bool, 0)
			for _, c := range prevLine {
				if _, ok := firstChars[c]; ok {
					potentialCommon[c] = true
				}
			}
			fmt.Println(firstChars)
			var common rune
			for _, c := range line {
				if _, ok := potentialCommon[c]; ok {

					common = c
					break
				}
			}
			fmt.Println(potentialCommon)
			ascii := int(common)
			// fmt.Printf("COMMON %c\n", common)

			if ascii <= 90 {
				// UPPERCASE
				ascii = ascii - 64 + 26
			} else {
				ascii -= 96
				// LOWERCASE
			}
			// fmt.Printf("%c -> %d\n", common, ascii)
			total += ascii
			fmt.Println("~~~")
		}
		i++
		// take first and second parts,
		// find common letter with map of first
		// add priority of common to total
	}
	fmt.Printf("Total %d\n", total)
	readFile.Close()
}
