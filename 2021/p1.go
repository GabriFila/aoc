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

	last := -1
	inc := 0
	prev2 := -1
	prev1 := -1
	cur := -1
	fmt.Printf("---------------\n\n\n")
	for fileScanner.Scan() {

		line := fileScanner.Text()

		if res, err := strconv.Atoi(line); err != nil {
			fmt.Printf("ERROR, %s\n", err)
			break
		} else {
			if prev1 != -1 && prev2 != -1 {

				cur = res + prev1 + prev2
				fmt.Printf("%d -> %d %d\n", res, cur, last)

				if last != -1 && cur > last {
					fmt.Printf("YES \n")
					inc += 1
				} else {

					fmt.Printf("NO \n")
				}
				last = cur
			}
			prev2 = prev1
			prev1 = res
			// fmt.Printf("Res %d %d %d\n", res, prev1, prev2)
		}
	}
	// cur = prev1 + prev1 + prev2
	// fmt.Printf("%d -> %d %d\n", prev1, cur, last)
	// // fmt.Printf("Res %d %d %d\n", prev1, prev1, prev2)
	// if cur > last {
	// 	fmt.Printf("YES \n")
	// 	inc++
	// } else {

	// 	fmt.Printf("NO \n")
	// }

	// last = cur
	// cur = prev1 + prev1 + prev1
	// fmt.Printf("%d -> %d %d\n", prev1, cur, last)
	// // fmt.Printf("Res %d %d %d\n", prev1, prev1, prev1)
	// if cur > last {
	// 	fmt.Printf("YES \n")
	// 	inc++
	// } else {
	// 	fmt.Printf("NO \n")

	// }

	fmt.Printf("Total %d\n", inc)
	readFile.Close()
}
