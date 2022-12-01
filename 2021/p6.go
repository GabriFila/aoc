package main

import (
	"fmt"
	"math"
)

const WIDTH = 1000
const HEIGHT = 1000

type State []int64

func (s State) print() {

	for _, val := range s {
		fmt.Print(val)
		fmt.Print(",")
	}
	fmt.Print("\n")
}

func childrenOfFish(days int64, init int64, level int) int64 {
	base := int64(0)

	fmt.Printf("~~ %d\n", level)
	fmt.Printf("Days: %d\n", days)
	fmt.Printf("Init: %d\n", init)
	if init > 0 {
		base = 1
		fmt.Printf("Base %d\n", base)
	}

	actualDays := days - init
	fmt.Printf("Actual: %d\n", actualDays)

	temp := int64(math.Floor(float64(actualDays) / 7))
	fmt.Printf("Temp: %d\n", temp)

	if init > 0 {
		// temp += childrenOfFish(days-init-2, 0, level+1)
	}

	if temp <= 0 {
		return 0
	} else {
		return temp
		// return temp + childrenOfFish(actualDays-7-2, 0, level+1)
	}
}

func main() {

	state := State{1}
	// state := State{3, 4, 3, 1, 2}
	// state := State{2, 5, 5, 3, 2, 2, 5, 1, 4, 5, 2, 1, 5, 5, 1, 2, 3, 3, 4, 1, 4, 1, 4, 4, 2, 1, 5, 5, 3, 5, 4, 3, 4, 1, 5, 4, 1, 5, 5, 5, 4, 3, 1, 2, 1, 5, 1, 4, 4, 1, 4, 1, 3, 1, 1, 1, 3, 1, 1, 2, 1, 3, 1, 1, 1, 2, 3, 5, 5, 3, 2, 3, 3, 2, 2, 1, 3, 1, 3, 1, 5, 5, 1, 2, 3, 2, 1, 1, 2, 1, 2, 1, 2, 2, 1, 3, 5, 4, 3, 3, 2, 2, 3, 1, 4, 2, 2, 1, 3, 4, 5, 4, 2, 5, 4, 1, 2, 1, 3, 5, 3, 3, 5, 4, 1, 1, 5, 2, 4, 4, 1, 2, 2, 5, 5, 3, 1, 2, 4, 3, 3, 1, 4, 2, 5, 1, 5, 1, 2, 1, 1, 1, 1, 3, 5, 5, 1, 5, 5, 1, 2, 2, 1, 2, 1, 2, 1, 2, 1, 4, 5, 1, 2, 4, 3, 3, 3, 1, 5, 3, 2, 2, 1, 4, 2, 4, 2, 3, 2, 5, 1, 5, 1, 1, 1, 3, 1, 1, 3, 5, 4, 2, 5, 3, 2, 2, 1, 4, 5, 1, 3, 2, 5, 1, 2, 1, 4, 1, 5, 5, 1, 2, 2, 1, 2, 4, 5, 3, 3, 1, 4, 4, 3, 1, 4, 2, 4, 4, 3, 4, 1, 4, 5, 3, 1, 4, 2, 2, 3, 4, 4, 4, 1, 4, 3, 1, 3, 4, 5, 1, 5, 4, 4, 4, 5, 5, 5, 2, 1, 3, 4, 3, 2, 5, 3, 1, 3, 2, 2, 3, 1, 4, 5, 3, 5, 5, 3, 2, 3, 1, 2, 5, 2, 1, 3, 1, 1, 1, 5, 1}
	fmt.Print("------------------\n\n\n")

	var days int64 = 28

	totalChildren := int64(0)
	for i, init := range state {
		children := childrenOfFish(days, init, 0)
		fmt.Printf("%d -> %d\n", i, children)
		totalChildren += children
	}

	fmt.Printf("Total: %d\n", totalChildren+int64(len(state)))

	for i := int64(0); i < days; i++ {
		curFishes := len(state)
		for i, val := range state {
			if i == curFishes {
				break
			}
			if val == 0 {
				state = append(state, 8)
				state[i] = 6
			} else {
				state[i]--
			}
		}
	}
	fmt.Printf("Total: %d\n", len(state))
}
