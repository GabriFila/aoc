package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var fileRgx = regexp.MustCompile(`(?m)([0-9]+)[ ]{1}`)

func strToint(in string) int {

	if ret, err := strconv.Atoi(in); err != nil {
		fmt.Printf("ERROR, %s\n", err)
		os.Exit(1)
	} else {
		return ret
	}
	return -1
}

func main() {

	path := make([]string, 0)
	readFile, err := os.Open("input.txt")

	dirSizes := make(map[string]int)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	fmt.Print("------------\n\n\n")

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Printf("Line: %s \n", line)
		if line[0] == '$' {
			if line[2:4] == "cd" {
				if line[5] == '/' {
					// fmt.Println("ROOT")
					path = []string{"/"}
				} else if line[5] == '.' {
					// fmt.Println("..")
					path = path[0 : len(path)-1]
				} else {
					// fmt.Println("Dir", line[5:])
					path = append(path, line[5:])
				}
			}
		} else {
			size := 0
			for i, match := range fileRgx.FindStringSubmatch(line) {
				if i == 1 {
					// fmt.Printf("Size: %s ..\n", match)
					size = strToint(match)
				}
			}

			if size > 0 {
				for i := range path {
					// fmt.Printf("%v\n", path[0:i+1])
					dirSizes[strings.Join(path[0:i+1], "~")] += size
				}
			}
		}
		// fmt.Println("Path", path)
		// fmt.Println("Sizes", dirSizes)
		// fmt.Println("~~~~")
	}

	total := 0

	min := 70000000
	used := dirSizes["/"]
	unsued := min - used
	toFree := 30000000 - unsued
	for _, size := range dirSizes {
		if size <= 100000 {
			// fmt.Printf("Size %d\n", size)
			total += size
		}
		if size >= toFree && size < min {
			min = size
		}
	}
	readFile.Close()

	fmt.Printf("Total %d\n", total)
	fmt.Println("/")
	fmt.Println("Unused", unsued)
	fmt.Println("Tofree", toFree)
	fmt.Println("Min", min)
}
