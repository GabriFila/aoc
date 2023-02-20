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

type Pair []string
type Pairs []Pair

func (p Pair) String() string {

	ret := ""

	for _, val := range p {
		ret += val
		ret += "\n"
	}
	return ret
}

func (p Pairs) String() string {

	ret := ""

	for i, pair := range p {
		ret += fmt.Sprintf("%v", pair)
		if i != len(p)-1 {
			ret += "\n"
		}
	}
	return ret
}

func compareValues(left string, right string, cursor int) (correctOrder bool) {
	if left[cursor] == '[' && right[0] == '[' {
		// both list

		
		return true
	} else if left[cursor] != '[' && right[0] != '[' {
		// both ints
		nextComma := strings.Index(left, ",")
		valL := strToInt(left[cursor:nextComma])
		valR := strToInt(right[cursor:nextComma])
		if valL < valR {
			return true
		} else if valR < valR {
			return false
		} else {
			// next elm
			return compareValues(left, right, cursor+2)
		}

	} else {
		return false
	}
}

func (p Pair) isOrderCorrect() bool {
	left := p[0]
	right := p[1]
	for cursor := 0; cursor < len(p[0]); cursor++ {
		if left[cursor] == '[' && right[0] == '[' {
			// both list
			return true
		} else if left[cursor] != '[' && right[0] != '[' {
			// both ints
			nextComma := strings.Index(left, ",")
			valL := strToInt(left[cursor:nextComma])
			valR := strToInt(right[cursor:nextComma])
			if valL < valR {
				return true
			} else if valR < valR {
				return false
			} else {
				// next elm
			}

		} else {

		}
	}
	return false
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	fmt.Print("------------\n\n\n")

	pairs := make(Pairs, 0)
	i := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line != "" {
			if i%2 == 0 {
				pairs = append(pairs, []string{})
				pairs[i/2] = make([]string, 0)
				pairs[i/2] = append(pairs[i/2], line)
			} else {
				pairs[i/2] = append(pairs[i/2], line)

			}
			i++
		}
	}

	for _, pair := range pairs {
		fmt.Println(pair, pair.isOrderCorrect())

	}
	// fmt.Println(pairs)
	total := 0
	fmt.Printf("Total %d\n", total)
	readFile.Close()
}
