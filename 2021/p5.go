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

const WIDTH = 1000
const HEIGHT = 1000

type Grid [][]int

func (g Grid) print() {

	if len(g) > 10 {
		return
	}
	fmt.Print(g)
	fmt.Print("\n")
	for _, row := range g {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func main() {
	reg := regexp.MustCompile(`[" "]+`)
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fmt.Printf("---------------\n\n\n")

	grid := make(Grid, HEIGHT)
	for i := range grid {
		grid[i] = make([]int, WIDTH)
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		startEnd := strings.Split(reg.ReplaceAllString(line, ""), "->")
		fmt.Println(startEnd)

		startXY := strings.Split(startEnd[0], ",")
		endXY := strings.Split(startEnd[1], ",")
		// fmt.Println(startXY)
		// fmt.Println(endXY)
		if startX, err := strconv.Atoi(startXY[0]); err != nil {
			fmt.Printf("ERROR, %s\n", err)
			os.Exit(1)
			break
		} else {
			if startY, err := strconv.Atoi(startXY[1]); err != nil {
				fmt.Printf("ERROR, %s\n", err)
				os.Exit(1)
				break
			} else {
				if endX, err := strconv.Atoi(endXY[0]); err != nil {
					fmt.Printf("ERROR, %s\n", err)
					os.Exit(1)
					break
				} else {
					if endY, err := strconv.Atoi(endXY[1]); err != nil {
						fmt.Printf("ERROR, %s\n", err)
						os.Exit(1)
						break
					} else {
						fmt.Printf("%d %d %d %d\n", startX, startY, endX, endY)

						if startX == endX || startY == endY {

							if startX > endX {
								temp := startX
								startX = endX
								endX = temp
							}
							if startY > endY {
								temp := startY
								startY = endY
								endY = temp
							}
						}
						if startX == endX {
							for y := startY; y < endY+1; y++ {
								grid[y][startX]++
								// fmt.Printf("Marking %d\n", grid[y][startX])
							}

						} else if startY == endY {
							for x := startX; x < endX+1; x++ {
								grid[startY][x]++
								// fmt.Printf("Marking %d\n", grid[startY][x])
							}
						} else {
							len := int(math.Abs(float64(endY-startY))) + 1
							fmt.Printf("Len %d\n", len)

							xSign := 1
							ySign := 1

							if endX < startX {
								xSign = -1
							}
							if endY < startY {
								ySign = -1
							}

							fmt.Printf("X %d %d\n", xSign, startX)
							fmt.Printf("Y %d %d\n", ySign, startY)
							for i := 0; i < len; i++ {
								targetX := startX + i*xSign
								targetY := startY + i*ySign
								// fmt.Printf("Marking at %d,%d -> %d\n", targetX, targetY, grid[targetY][targetX])
								grid[targetY][targetX] += 1
							}

						}
					}
				}
			}
		}

		grid.print()
	}
	total := 0

	// grid.print()
	for _, row := range grid {
		for _, val := range row {
			if val >= 2 {
				total++
			}

		}
	}
	fmt.Printf("Total %d\n", total)

	readFile.Close()
}
