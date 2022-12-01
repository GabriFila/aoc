package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var amountRgx = regexp.MustCompile(`(?m)move (\d+) from`)
var fromRgx = regexp.MustCompile(`(?m)from (\d+) to`)
var toRgx = regexp.MustCompile(`(?m)to (\d+)`)

func byteToInt(in byte) int {

	if ret, err := strconv.Atoi(string(in)); err != nil {
		fmt.Printf("ERROR, %s\n", err)
		os.Exit(1)
	} else {
		return ret
	}
	return -1
}
func strToInt(in string) int {

	if ret, err := strconv.Atoi(in); err != nil {
		fmt.Printf("ERROR, %s\n", err)
		os.Exit(1)
	} else {
		return ret
	}
	return -1
}

func computeDistance(hX, hY, tX, tY int) int {
	if hX == tX {
		return int(math.Abs(float64(tY) - float64(hY)))
	} else if hY == tY {
		return int(math.Abs(float64(tX) - float64(hX)))
	} else {
		return int(math.Sqrt(float64(hX)*float64(tX) + float64(hY)*float64(tY)))
	}
}

type Passes map[int]map[int]bool

func (p Passes) set(x, y int) {
	if _, ok := p[x]; !ok {
		p[x] = map[int]bool{}
	}
	p[x][y] = true
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	fmt.Println("------------\n\n\n")

	lines := make([]string, 0)
	tX := 0
	tY := 0

	alreadyPassed := make(Passes, 0)
	alreadyPassed.set(tX, tY)

	x := 1
	cycle := 1
	totStrength := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)

		if line == "noop" {
			// fmt.Printf("Cycle %d NOOP\n", cycle)
		} else {
			cycle++

			if (cycle-20)%40 == 0 {
				fmt.Printf("Cycles %d\n", cycle)

				totStrength += x * cycle
			}
			amount := strToInt(line[5:])
			x += amount
		}
		cycle++
		if (cycle-20)%40 == 0 {
			fmt.Printf("Cycles %d\n", cycle)
			totStrength += x * cycle
		}
	}

	fmt.Printf("Total %d\n", totStrength)
	fmt.Printf("Cycles %d\n", cycle)
	fmt.Printf("X %d\n", x)
	readFile.Close()
}
