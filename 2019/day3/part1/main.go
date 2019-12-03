package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var inputInts []int

func main() {
	t := time.Now()
	inputFile := os.Args[1]
	inputList := readInFile(inputFile)
	lineOne := strings.Split(inputList[0], ",")
	lineTwo := strings.Split(inputList[1], ",")
	lineOnePoints := drawLine(lineOne)
	lineTwoPoints := drawLine(lineTwo)
	intersections := []string{}
	for k := range lineTwoPoints {
		if _, ok := lineOnePoints[k]; ok {
			intersections = append(intersections, k)
		}
	}
	//for k := range lineOnePoints {
	//	fmt.Println(k)
	//}
	manhattenPoint := 10000
	manhattenX := 0
	manhattenY := 0
	for _, pointString := range intersections {
		if pointString != "0,0" {
			points := strings.Split(pointString, ",")
			xPoint, _ := strconv.Atoi(points[0])
			yPoint, _ := strconv.Atoi(points[1])
			absX := abs(xPoint)
			absY := abs(yPoint)
			manPoint := absX + absY
			if manPoint < manhattenPoint {
				manhattenPoint = manPoint
				manhattenX = xPoint
				manhattenY = yPoint
			}
		}
	}
	fmt.Println("ManhattenPoint: ", manhattenPoint)
	fmt.Println("orig x: ", manhattenX, "orig y: ", manhattenY)
	fmt.Println(time.Since(t))
}

func drawLine(instructionSet []string) map[string]int {
	x := 0
	y := 0
	count := 0
	workingArray := map[string]int{"0,0": 1}
	for _, instructionBlock := range instructionSet {
		instruction := strings.Split(instructionBlock, "")
		switch instruction[0] {
		case "R":
			tempX, _ := strconv.Atoi(strings.Join(instruction[1:], ""))
			for i := 1; i <= tempX; i++ {
				point := fmt.Sprintf("%d,%d", i+x, y)
				workingArray[point] = count
			}
			x += tempX
		case "L":
			tempX, _ := strconv.Atoi(strings.Join(instruction[1:], ""))
			for i := -1; i >= -tempX; i-- {
				point := fmt.Sprintf("%d,%d", i+x, y)
				workingArray[point] = count
			}
			x += -tempX
		case "U":
			tempY, _ := strconv.Atoi(strings.Join(instruction[1:], ""))
			for i := 1; i <= tempY; i++ {
				point := fmt.Sprintf("%d,%d", x, i+y)
				workingArray[point] = count
			}
			y += tempY
		case "D":
			tempY, _ := strconv.Atoi(strings.Join(instruction[1:], ""))
			for i := -1; i >= -tempY; i-- {
				point := fmt.Sprintf("%d,%d", x, i+y)
				workingArray[point] = count
			}
			y += -tempY
		}
	}
	return workingArray

}

func abs(point int) (absPoint int) {
	if point > 0 {
		return point
	}
	return -point
}
func readInFile(fileName string) []string {
	input := []string{}
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return input
}
