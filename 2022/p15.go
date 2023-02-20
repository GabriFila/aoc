package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func strToInt(in string) int {

	if ret, err := strconv.Atoi(in); err != nil {
		fmt.Printf("ERROR, %s\n", err)
		os.Exit(1)
		return 0
	} else {
		return int(ret)
	}
}

type OperationType byte

const (
	Add OperationType = '+'
	Mul OperationType = '*'
)

type Monkey struct {
	startingItems []int
	opType        OperationType
	opVal         int
	trueNext      int
	falseNext     int
	divBy         int
	inspected     int
}

var xRgx = regexp.MustCompile(`(?m)x=(.*),`)
var yRgx = regexp.MustCompile(`(?m)y=(.*)`)

func getFirstRgxGroup(input string, rgx *regexp.Regexp) string {
	for _, match := range rgx.FindAllStringSubmatch(input, -1) {
		return match[1]
	}
	return ""
}

type Line struct {
	xSensor int
	ySensor int
	xBeacon int
	yBeacon int
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	fmt.Print("------------\n\n\n")

	// compute lowest and highest X (xMin, xMax)
	// fill array of length abs(xMax - xMin) with false
	// take sensor, compute y extension, check if target line is reached by the sensor
	// if no skip
	// if yes
	// compute X distance between S and B -> x = abs(Xsensor-Xbeacon)
	// check distance D between sensor and target line
	// compute R range at target line by doing D-X
	// file arr with true between xSensor-R and xSensor+R+1
	lines := make([]Line, 0)
xMin := 0
xMax
	for fileScanner.Scan() {
		line := fileScanner.Text()

		dataStr := strings.Split(line, ":")

		sensorStr := dataStr[0]
		beaconStr := dataStr[1]
		xSensor := strToInt(getFirstRgxGroup(sensorStr, xRgx))
		xBeacon := strToInt(getFirstRgxGroup(beaconStr, xRgx))

		lines = append(lines, Line{
			xSensor: xSensor,
			ySensor: strToInt(getFirstRgxGroup(sensorStr, yRgx)),
			xBeacon: xBeacon,
			yBeacon: strToInt(getFirstRgxGroup(beaconStr, yRgx)),
		})
		readFile.Close()
	}
	// total := 0
	// fmt.Printf("Total %d\n", total)
}
