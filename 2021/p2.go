package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	depth := 0
	forward := 0
	aim := 0
	fmt.Printf("---------------\n\n\n")

	for fileScanner.Scan() {

		line := fileScanner.Text()
		if strings.HasPrefix(line, "forward") {
			if res, err := strconv.Atoi(string(line[8])); err != nil {
				fmt.Printf("ERROR, %s\n", err)
				break
			} else {

				forward += res
				depth += aim * res
			}
		} else if strings.HasPrefix(line, "down") {
			if res, err := strconv.Atoi(string(line[5])); err != nil {
				fmt.Printf("ERROR, %s\n", err)
				break
			} else {
				aim += res
			}

		} else {
			if res, err := strconv.Atoi(string(line[3])); err != nil {
				fmt.Printf("ERROR, %s\n", err)
				break
			} else {
				aim -= res
			}

		}

	}

	fmt.Printf("Total %d\n", depth*forward)
	readFile.Close()
}
