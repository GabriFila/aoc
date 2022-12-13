package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	fmt.Print("------------\n\n\n")

	lines := make([]string, 0)
	hX := 0
	hY := 0
	tX := 0
	tY := 0
	total := 0

	alreadyPassed := make(Passes, 0)
	alreadyPassed.set(tX, tY)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)

		lineSplit := strings.Split(line, " ")

		dir := lineSplit[0]
		amount := strToInt(lineSplit[1])
		fmt.Printf("Dir %s\n", dir)
		if dir == "R" {
			for i := 0; i < amount; i++ {
				hX++
				if hX == tX && hY == tY {
					println("Overlap")
				} else if hX == tX {
					if int(math.Abs(float64(tY)-float64(hY))) > 1 {
						println("Different row")
						tY++
					}
				} else if hY == tY {
					if int(math.Abs(float64(tX)-float64(hX))) > 1 {
						println("Different col")
						tX++
					}
				} else {
					if int(math.Sqrt(float64(hX)*float64(tX)+float64(hY)*float64(tY))) > 1 {
						println("Diagonal")
						tX++
					}
				}
				alreadyPassed.set(tX, tY)
			}
		}
		break
	}

	fmt.Printf("hX: %d,hY: %d,tX: %d,tY: %d\n", hX, hY, tX, tY)
	fmt.Printf("Passes %v\n", alreadyPassed)
	fmt.Printf("Total %d\n", total)
	readFile.Close()
}
