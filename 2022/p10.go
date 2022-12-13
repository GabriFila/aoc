package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	fmt.Print("------------\n\n\n")

	x := 1
	cycle := 1
	totStrength := 0

	display := ""
	for fileScanner.Scan() {
		fmt.Printf("Sprite %d\n", x)
		line := fileScanner.Text()

		cursor := (cycle - 1) % 40
		if cursor == 0 {
			display += "\n"
		}
		if cursor <= x+1 && cursor >= x-1 {
			display += "#"
		} else {
			display += "."
		}
		cycle++

		if line == "noop" {
			fmt.Printf("Cycle %d NOOP\n", cycle)
		} else {

			fmt.Printf("Cycle %d\n", cycle)
			cursor := (cycle - 1) % 40
			fmt.Printf("Cursor %d\n", cursor)
			if cursor == 0 {
				display += "\n"
			}
			if cursor <= x+1 && cursor >= x-1 {
				display += "#"
			} else {
				display += "."
			}
			fmt.Println(display)
			if (cycle-20)%40 == 0 {
				// fmt.Printf("Cycles %d\n", cycle)

				totStrength += x * cycle
			}
			cycle++
			amount := strToInt(line[5:])
			x += amount
			fmt.Printf("Sprite %d\n", x)
		}

		if (cycle-20)%40 == 0 {
			// fmt.Printf("Cycles %d\n", cycle)
			totStrength += x * cycle
		}
		fmt.Println(display)
		fmt.Println("~~~")
	}

	fmt.Printf("Total %d\n", totStrength)
	fmt.Printf("Cycles %d\n", cycle)
	fmt.Printf("X %d\n", x)
	readFile.Close()
}
