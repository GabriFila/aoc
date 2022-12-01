package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var myMovePoints = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	// var oppMovePoints = map[string]int{
	// 	"A": 1,
	// 	"B": 2,
	// 	"C": 3,
	// }
	var myMoveForOutcome = map[string]map[string]string{
		"X": {
			"A": "Z",
			"B": "X",
			"C": "Y",
		},
		"Y": {
			"A": "X",
			"B": "Y",
			"C": "Z",
		},
		"Z": {
			"A": "Y",
			"B": "Z",
			"C": "X",
		},
	}

	total := 0
	fmt.Print("------------\n\n\n")
	for fileScanner.Scan() {

		line := fileScanner.Text()

		oppMove := string(line[0])
		myMove := string(line[2])

		// myPts := myMovePoints[myMove]
		// oppPts := oppMovePoints[oppMove]

		// total += myPts
		// fmt.Printf("opp pts %d\n", oppPts)
		// fmt.Printf("my %s %d\n", myMove, myPts)
		// fmt.Printf("opp %s %d\n", oppMove, oppPts)
		switch myMove {
		case "X":
			total += 0
			fmt.Println("LOST")
			break
		case "Y":
			fmt.Println("DRAW")
			total += 3
			break
		case "Z":
			fmt.Println("WIN")
			total += 6
			break
		}

		myMovePts := myMovePoints[myMoveForOutcome[myMove][oppMove]]
		total += myMovePts
		fmt.Printf("my %s->%d\n", myMoveForOutcome[myMove][oppMove], myMovePts)

		// if myPts == oppPts {
		// 	fmt.Println("DRAW")
		// 	total += 3
		// } else if myMove == "X" && oppMove == "C" || myMove == "Y" && oppMove == "A" || myMove == "Z" && oppMove == "B" {
		// 	total += 6
		// 	fmt.Println("WON")

		// } else {
		// 	fmt.Println("LOST")
		// }
	}
	fmt.Printf("Total %d\n", total)
	readFile.Close()
}
