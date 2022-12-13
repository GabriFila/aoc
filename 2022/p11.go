package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strToInt(in string) uint64 {

	if ret, err := strconv.Atoi(in); err != nil {
		fmt.Printf("ERROR, %s\n", err)
		os.Exit(1)
		return 0
	} else {
		return uint64(ret)
	}
}

type OperationType byte

const (
	Add OperationType = '+'
	Mul OperationType = '*'
)

type Monkey struct {
	startingItems []uint64
	opType        OperationType
	opVal         uint64
	trueNext      uint64
	falseNext     uint64
	divBy         uint64
	inspected     uint64
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	fmt.Print("------------\n\n\n")

	monkeys := make([]Monkey, 0)
	i := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			i++
		} else if line[0] == 'M' {
			monkeys = append(monkeys, Monkey{})
		} else if line[2] == 'S' {
			elms := strings.Split(line[18:], ", ")
			for _, val := range elms {

				monkeys[i].startingItems = append(monkeys[i].startingItems, strToInt(val))
			}

		} else if line[2] == 'O' {
			monkeys[i].opType = OperationType(line[23])
			val := line[25:]
			if val == "old" {
				monkeys[i].opVal = 0
			} else {
				monkeys[i].opVal = strToInt(val)
			}
		} else if line[2] == 'T' {
			monkeys[i].divBy = strToInt(line[21:])
		} else if line[7] == 't' {
			monkeys[i].trueNext = strToInt(line[29:])
		} else if line[7] == 'f' {
			monkeys[i].falseNext = strToInt(line[30:])

		}

		// fmt.Println("~~~")
	}

	for _, monkey := range monkeys {
		fmt.Printf("%+v\n\n", monkey)
	}
	rounds := 10000
	for i := 0; i < rounds; i++ {
		for monkeyIdx := range monkeys {
			items := len(monkeys[monkeyIdx].startingItems)
			monkeys[monkeyIdx].inspected += uint64(items)
			for itemIdx := 0; itemIdx < items; itemIdx++ {
				worryLevel := monkeys[monkeyIdx].startingItems[0]
				// fmt.Printf("Round %d, monkey %d, items %d\n", i, monkeyIdx, worryLevel)

				if monkeys[monkeyIdx].opType == Mul {
					if monkeys[monkeyIdx].opVal == 0 {
						worryLevel *= worryLevel
						fmt.Println("BAZINGA", worryLevel)
					} else {
						worryLevel *= monkeys[monkeyIdx].opVal
					}
				} else {
					if monkeys[monkeyIdx].opVal == 0 {
						worryLevel += worryLevel
					} else {
						worryLevel += monkeys[monkeyIdx].opVal
					}
				}
				// worryLevel /= 3
				if worryLevel%monkeys[monkeyIdx].divBy == 0 {
					monkeys[monkeys[monkeyIdx].trueNext].startingItems = append(monkeys[monkeys[monkeyIdx].trueNext].startingItems, worryLevel)
				} else {
					monkeys[monkeys[monkeyIdx].falseNext].startingItems = append(monkeys[monkeys[monkeyIdx].falseNext].startingItems, worryLevel)
				}
				// fmt.Println("Bef Start", monkeys[monkeyIdx].startingItems)
				newStartingItems := monkeys[monkeyIdx].startingItems[1:]

				// fmt.Println("Aft Start", newStartingItems)
				monkeys[monkeyIdx].startingItems = newStartingItems
				// fmt.Println("BUL")
			}
		}

		if i == 19 {
			fmt.Printf("ROUND %d\n", i)
			for i := range monkeys {

				fmt.Printf("Monkey %d: %d\n", i, monkeys[i].inspected)
				// fmt.Printf("%+v\n\n", monkey)
			}
			break
		}
	}
	fmt.Println()

	mostTimes := []uint64{0, 0}
	minIdx := 0
	for i, monkey := range monkeys {
		if monkey.inspected > mostTimes[minIdx] {
			mostTimes[minIdx] = monkey.inspected
		}
		newMin := mostTimes[0]
		minIdx = 0
		for j := range mostTimes {
			if mostTimes[j] < newMin {
				minIdx = j
			}
		}

		fmt.Printf("Monkey %d: %d\n", i, monkey.inspected)
		// fmt.Printf("%+v\n\n", monkey)
	}
	fmt.Printf("Most times %d\n", mostTimes)

	total := uint64(1)

	for _, times := range mostTimes {
		total *= times
	}
	fmt.Printf("Total %d\n", total)
	readFile.Close()
}
