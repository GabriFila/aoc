package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func strToint(in string) int {

	if ret, err := strconv.Atoi(in); err != nil {
		fmt.Printf("ERROR, %s\n", err)
		os.Exit(1)
	} else {
		return ret
	}
	return -1
}

func remove(slice []byte, at int) []byte {
	return append(slice[:at], slice[at+1:]...)
}

type Stacks []Stack
type Stack []byte

func (s Stack) String() string {
	ret := "["
	length := len(s)
	for i, val := range s {
		ret += string(val)
		if i != length-1 {
			ret += ","
		}
	}
	ret += "]"
	return ret
}
func (s Stacks) String() string {
	ret := ""
	for i, val := range s {
		ret += fmt.Sprintf("%d %s\n", i, val)
	}
	return ret
}

var amountRgx = regexp.MustCompile(`(?m)move (\d+) from`)
var fromRgx = regexp.MustCompile(`(?m)from (\d+) to`)
var toRgx = regexp.MustCompile(`(?m)to (\d+)`)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	stacks := make(Stacks, 9)
	for i := range stacks {
		stacks[i] = make(Stack, 0)
	}

	total := 0
	fmt.Print("------------\n\n\n")

	isFillingStacks := true
	for fileScanner.Scan() {

		line := fileScanner.Text()
		// fmt.Printf("Line %s\n", line)

		if isFillingStacks {

			if string(line[1]) == "1" {
				fmt.Printf("AAAAAAAAAAAAA\n")
				isFillingStacks = false
			} else {

				for i := range stacks {
					idx := 1 + i*4
					char := line[idx]
					// fmt.Printf("Char at %d %c %v\n", idx, char, rune(char))
					if rune(char) != 32 {
						// fmt.Printf("Att %c\n", char)
						// fmt.Printf("Stacks %v\n", stacks)
						stacks[i] = append(Stack{line[idx]}, stacks[i]...)
					}

				}
			}
		} else if len(line) != 0 {
			amount := 0

			for i, match := range amountRgx.FindStringSubmatch(line) {
				if i == 1 {
					fmt.Printf("Amount: %s ..\n", match)
					amount = strToint(match)
				}
			}

			from := -1
			for i, match := range fromRgx.FindStringSubmatch(line) {
				if i == 1 {
					fmt.Printf("From: %s ..\n", match)
					from = strToint(match) - 1
				}
			}
			to := -1
			for i, match := range toRgx.FindStringSubmatch(line) {
				if i == 1 {
					fmt.Printf("To: %s ..\n", match)
					to = strToint(match) - 1
				}
			}

			// fmt.Printf("IsFilling stacks %t\n", isFillingStacks)

			for i := 0; i < amount; i++ {
				fmt.Printf("I: %d \n", i)
				movedIdx := len(stacks[from]) - (amount) + i
				fmt.Printf("Len: %d \n", movedIdx)

				movedElm := stacks[from][movedIdx]
				// fmt.Printf("Moved %c\n", movedElm)
				stacks[from] = remove(stacks[from], movedIdx)
				// fmt.Printf("From %v\n", stacks[from])
				// fmt.Printf("From %v\n", stacks[from])
				// fmt.Printf("Moved %c\n", movedElm)
				stacks[to] = append(stacks[to], movedElm)
				// fmt.Printf("To %v\n", stacks[to])
			}
			fmt.Println(stacks)
		}

	}
	// fmt.Printf("Stacks %v\n", stacks)
	ret := ""
	for i := range stacks {
		ret += fmt.Sprintf("%c", stacks[i][len(stacks[i])-1])
	}
	fmt.Printf("Total %d\n", total)
	fmt.Printf("Ret %s\n", ret)
	readFile.Close()
}
