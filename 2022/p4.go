package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task string

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func strToint(in string) int {

	if ret, err := strconv.Atoi(in); err != nil {
		fmt.Printf("ERROR, %s\n", err)
		os.Exit(1)
	} else {
		return ret
	}
	return -1
}

func (t Task) String() string {
	startA, endA, startB, endB := t.getPoints()
	first := ".........."
	second := ".........."

	for i := startA; i < endA+1; i++ {
		var temp rune
		for _, c := range fmt.Sprintf("%d", i) {
			temp = c
			break
		}
		first = replaceAtIndex(first, temp, i)
	}
	for i := startB; i < endB+1; i++ {
		var temp rune
		for _, c := range fmt.Sprintf("%d", i) {
			temp = c
			break
		}
		second = replaceAtIndex(second, temp, i)
	}
	ret := ""
	ret += fmt.Sprintf("%s\n", first)
	ret += fmt.Sprintf("%s", second)
	return ret
}
func (t Task) getPoints() (int, int, int, int) {
	pairs := strings.Split(string(t), ",")

	pairA := strings.Split(pairs[0], "-")
	pairB := strings.Split(pairs[1], "-")

	return toInt(pairA[0]), toInt(pairA[1]), toInt(pairB[0]), toInt(pairB[1])
}
func (t Task) isOverlap() {

}

func toInt(str string) int {
	if res, err := strconv.Atoi(str); err != nil {
		fmt.Printf("ERROR, %s\n", err)
		panic(0)
	} else {
		return res
	}
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	total := 0
	fmt.Print("------------\n\n\n")

	line := ""

	for fileScanner.Scan() {
		line = fileScanner.Text()

		fmt.Println(Task(line)) // compute
		startA, endA, startB, endB := Task(line).getPoints()
		// part 1
		// if startA <= startB && endA >= endB {
		// 	total++
		// } else if startB <= startA && endB >= endA {
		// 	total++
		// }
		// part 2
		if startA < startB && endA < startB || startA > endB && endA > endB {
			fmt.Println("NO OVERLAP")
		} else {
			total++
		}

		fmt.Println("~~~")
	}
	fmt.Printf("Total %d\n", total)
	readFile.Close()

}
