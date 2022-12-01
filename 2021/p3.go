package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fmt.Printf("---------------\n\n\n")
	// gamma := ""
	// epsilon := ""
	// totalZeros := make([]int, bits)

	lines := make([]string, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)

	}
	// fmt.Println("lines", lines)

	o2, err := findRating(lines, true)

	co2, err := findRating(lines, false)

	fmt.Println("RES", co2*o2)
	readFile.Close()
}

func findRating(lines []string, useMostCommon bool) (int64, error) {
	filtered := make([]string, len(lines))
	copy(filtered, lines)
	newFiltered := make([]string, 0)

	for i := 0; i < len(lines[0]); i++ {
		mostCommon := findCommon(filtered, useMostCommon, i)
		// fmt.Println("pos", i)
		// fmt.Println("mostcommon", mostCommon)

		for _, fil := range filtered {
			// fmt.Printf("at pos %d: %s\n", i, string(fil[i]))
			if string(fil[i]) == mostCommon {
				// fmt.Printf("added\n")
				newFiltered = append(newFiltered, fil)
			}
		}
		// fmt.Println("newFiltered", newFiltered)
		if len(newFiltered) == 1 {
			// fmt.Println(newFiltered)
			fmt.Println("FOUND")
			if res, err := strconv.ParseInt(newFiltered[0], 2, 64); err != nil {
				fmt.Println(err)
				return -1, err
			} else {
				return res, nil
			}
		} else {
			filtered = make([]string, len(newFiltered))
			copy(filtered, newFiltered)
			newFiltered = make([]string, 0)
		}
	}
	return -1, fmt.Errorf("COULD NOT FIND UNIQUE VALUE")
}

func findCommon(lines []string, useMostCommon bool, pos int) string {
	totalZeros := 0
	for _, line := range lines {
		if string(line[pos]) == "0" {
			totalZeros++
		}
	}

	if useMostCommon {
		if totalZeros > len(lines)/2 {
			return "0"
		} else {
			return "1"
		}
	} else {
		if totalZeros > len(lines)/2 {
			return "1"
		} else {
			return "0"
		}
	}
}
