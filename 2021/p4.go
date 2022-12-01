package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Board struct {
	data [][]int
}

func (r Board) won(numbers []int) (bool, int) {
	for _, row := range r.data {

		count := 0
		for _, b := range row {
			if contains(numbers, b) {
				count++
			} else {

				break
			}
		}
		if count == 5 {
			total := 0
			for _, row := range r.data {

				for _, b := range row {
					if !contains(numbers, b) {
						total += b
					}
				}
			}
			return true, total
		}
	}

	for i := 0; i < len(r.data); i++ {
		count := 0
		for _, row := range r.data {
			if contains(numbers, row[i]) {
				count++
			} else {
				break
			}
		}
		if count == 5 {

			total := 0
			for i := 0; i < len(r.data); i++ {
				for _, row := range r.data {
					if !contains(numbers, row[i]) {
						total += row[i]
					}
				}
			}
			return true, total
		}

	}
	return false, 0

}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fmt.Printf("---------------\n\n\n")

	numbers := []int{42, 44, 71, 26, 70, 92, 77, 45, 6, 18, 79, 54, 31, 34, 64, 32, 16, 55, 81, 11, 90, 10, 21, 87, 0, 84, 8, 23, 1, 12, 60, 20, 57, 68, 61, 82, 49, 59, 22, 2, 63, 33, 50, 39, 28, 30, 88, 41, 69, 72, 98, 73, 7, 65, 53, 35, 96, 67, 36, 4, 51, 75, 24, 86, 97, 85, 66, 29, 74, 40, 93, 58, 9, 62, 95, 91, 80, 99, 14, 19, 43, 37, 27, 56, 94, 25, 83, 48, 17, 38, 78, 15, 52, 76, 5, 13, 46, 89, 47, 3}
	// numbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	boards := make([]Board, 0)

	reg := regexp.MustCompile(`[" "]+`)

	lastRow := 0
	boardData := make([][]int, 5)
	for i := range boardData {
		boardData[i] = make([]int, 5)
	}
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			boards = append(boards, Board{data: boardData})
			boardData = make([][]int, 5)
			for i := range boardData {
				boardData[i] = make([]int, 5)
			}
			lastRow = 0
		} else {
			cleanedLine := reg.ReplaceAllString(line, " ")

			if string(cleanedLine[0]) == " " {
				cleanedLine = cleanedLine[1:]
			}

			numsInRow := strings.Split(cleanedLine, " ")
			for i, numInRow := range numsInRow {
				if res, err := strconv.Atoi(numInRow); err != nil {
					fmt.Printf("ERROR, %s\n", err)
					os.Exit(1)
					break
				} else {
					boardData[lastRow][i] = res
				}
			}
			lastRow++
		}

	}
	boards = append(boards, Board{data: boardData})

	fmt.Println("boards", boards)
	fmt.Println("boards", len(boards))

	// const MaxUint = ^uint(0)
	// const MinUint = 0
	// const MaxInt = int(MaxUint >> 1)
	// const MinInt = -MaxInt - 1
	lastWinStep := -1

	lastWinTotal := -1

	for _, board := range boards {
		fmt.Println("~~~~")
		fmt.Println("")
		fmt.Println("")
		for numIdx := range numbers {

			fmt.Println("id ", numIdx)
			fmt.Println("out ", numbers[numIdx])
			if won, tot := board.won(numbers[0 : numIdx+1]); won {

				// fmt.Print(boardIdx, " WON ")
				// fmt.Println("WITH ", numbers[numIdx])
				// fmt.Println("TOT ", tot, tot*numbers[numIdx])
				fmt.Println("tot", tot)
				if numIdx > lastWinStep {
					lastWinStep = numIdx
					lastWinTotal = tot * numbers[numIdx]
					fmt.Println("lastWinTotal", lastWinTotal)
				}
				break
			}

		}
	}
	fmt.Println("AT ", lastWinStep)
	fmt.Println("TOT ", lastWinTotal)

	readFile.Close()
}
