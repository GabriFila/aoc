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

	first := -1
	second := -1
	third := -1
	cur := 0
	for fileScanner.Scan() {

		line := fileScanner.Text()

		if line == "" {

			fmt.Printf("Empty %d\n", cur)
			if cur > first {
				third = second
				second = first
				first = cur
			}
			cur = 0
		} else {
			if res, err := strconv.Atoi(line); err != nil {
				fmt.Printf("ERROR, %s\n", err)
				break
			} else {
				cur += res
			}

			// fmt.Printf("Empty")

		}
	}

	fmt.Printf("First %d\n", first)
	fmt.Printf("Second %d\n", second)
	fmt.Printf("Third %d\n", third)
	fmt.Printf("Total %d\n", first+second+third)
	readFile.Close()
}
