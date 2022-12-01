package main

import (
	"bufio"
	"fmt"
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
func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	fmt.Println("------------\n\n\n")

	lines := make([]string, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	total := len(lines)*2 + 2*(len(lines[0])-2)
	fmt.Printf("Surrounding %d\n", total)

	maxScenicScore := 0

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			tree := byteToInt(lines[i][j])
			visibile := true

			// check top
			// for k := 0; k < j; k++ {
			for k := j - 1; k >= 0; k-- {
				if byteToInt(lines[i][k]) >= tree {
					visibile = false
					break
				}
			}

			if !visibile {
				visibile = true
				// check bottom
				// for k := j + 1; k < len(lines); k++ {
				for k := len(lines) - 1; k >= j+1; k-- {
					if byteToInt(lines[i][k]) >= tree {
						visibile = false
						break
					}
				}
				for k := len(lines) - 1; k >= j+1; k-- {
					if byteToInt(lines[i][k]) >= tree {
						visibile = false
						break
					}
				}

			}
			if !visibile {
				visibile = true
				// check left
				// for k := 0; k < i; k++ {
				for k := i - 1; k >= 0; k-- {
					if byteToInt(lines[k][j]) >= tree {
						visibile = false
						break
					}
				}

			}
			if !visibile {
				visibile = true
				// check left
				// for k := i + 1; k < len(lines[0]); k++ {
				for k := len(lines[0]) - 1; k >= i+1; k-- {
					if byteToInt(lines[k][j]) >= tree {
						visibile = false
						break
					}
				}

			}

			if visibile {
				total++
			}
			fmt.Printf("Elm %d visible: %t\n", tree, visibile)
		}
	}

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			tree := byteToInt(lines[i][j])
			leftScenicScore := 0
			for k := j - 1; k >= 0; k-- {
				cand := byteToInt(lines[i][k])
				if cand < tree {
					leftScenicScore++
				} else if cand >= tree {
					leftScenicScore++
					break
				}
			}

			topScenicScore := 0
			// check left
			// for k := 0; k < i; k++ {
			for k := i - 1; k >= 0; k-- {

				cand := byteToInt(lines[k][j])
				// fmt.Printf("BLBL %d\n", cand)
				if cand < tree {
					topScenicScore++
				} else if cand >= tree {
					topScenicScore++
					break
				}
			}

			rightScenicScore := 0
			// check bottom
			for k := j + 1; k < len(lines[0]); k++ {
				cand := byteToInt(lines[i][k])
				if cand < tree {
					rightScenicScore++
				} else if cand >= tree {
					rightScenicScore++
					break
				}

			}

			bottomScenicScore := 0
			// check right
			// for k := i + 1; k < len(lines[0]); k++ {
			fmt.Printf("i %d \n", i)
			fmt.Printf("j %d \n", j)
			for k := i + 1; k < len(lines); k++ {
				cand := byteToInt(lines[k][j])
				fmt.Printf("CAND %d \n", cand)
				if cand < tree {
					fmt.Printf("BL %d \n", cand)
					bottomScenicScore++
				} else if cand >= tree {
					fmt.Printf("BLU %d \n", cand)
					bottomScenicScore++
					break
				}
			}
			scenicScore := rightScenicScore * leftScenicScore * bottomScenicScore * topScenicScore
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
			fmt.Printf("Elm %d scenicScor: %d = %d*%d*%d*%d\n", tree, scenicScore, topScenicScore, leftScenicScore, bottomScenicScore, rightScenicScore)
		}
	}

	fmt.Printf("Total %d\n", total)
	fmt.Printf("Max Scenic score %d\n", maxScenicScore)
	readFile.Close()
}
