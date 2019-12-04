package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var inputInts []int

func main() {
	inputFile := os.Args[1]
	inputList := readInFile(inputFile)
	t := time.Now()
	lineOne := strings.Split(inputList[0], ",")
	lineTwo := strings.Split(inputList[1], ",")
	var listOne map[string]int
	var listTwo map[string]int
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(thisThing []string) {
		listOne = drawLine(thisThing)
		wg.Done()
	}(lineOne)

	go func(thisThing []string) {
		listTwo = drawLine(thisThing)
		wg.Done()
	}(lineTwo)
	wg.Wait()
	intersections := map[string]int{}
	for k := range listTwo {
		if _, ok := listOne[k]; ok {
			intersections[k] = listOne[k] + listTwo[k]
		}
	}
	lowSteps := 0
	manhattenX := 0
	manhattenY := 0
	for pointString, steps := range intersections {
		if lowSteps == 0 {
			lowSteps = steps
		} else {
			points := strings.Split(pointString, ",")
			xPoint, _ := strconv.Atoi(points[0])
			yPoint, _ := strconv.Atoi(points[1])
			if steps < lowSteps {
				lowSteps = steps
				manhattenX = xPoint
				manhattenY = yPoint
			}
		}
	}
	fmt.Println("Shortest Steps: ", lowSteps)
	fmt.Println("orig x: ", manhattenX, "orig y: ", manhattenY)
	fmt.Println(time.Since(t))
}

func drawLine(instructionSet []string) map[string]int {
	x := 0
	y := 0
	count := 1
	workingArray := map[string]int{"0,0": 0}
	for _, instructionBlock := range instructionSet {
		switch string(instructionBlock[0]) {
		case "R":
			tempX, _ := strconv.Atoi(instructionBlock[1:])
			for i := 1; i <= tempX; i++ {
				point := strconv.Itoa(i+x) + "," + strconv.Itoa(y)
				workingArray[point] = count
				count++
			}
			x += tempX
		case "L":
			tempX, _ := strconv.Atoi(instructionBlock[1:])
			for i := -1; i >= -tempX; i-- {
				point := strconv.Itoa(i+x) + "," + strconv.Itoa(y)
				workingArray[point] = count
				count++
			}
			x += -tempX
		case "U":
			tempY, _ := strconv.Atoi(instructionBlock[1:])
			for i := 1; i <= tempY; i++ {
				point := strconv.Itoa(x) + "," + strconv.Itoa(i+y)
				workingArray[point] = count
				count++
			}
			y += tempY
		case "D":
			tempY, _ := strconv.Atoi(instructionBlock[1:])
			for i := -1; i >= -tempY; i-- {
				point := strconv.Itoa(x) + "," + strconv.Itoa(i+y)
				workingArray[point] = count
				count++
			}
			y += -tempY
		}
	}
	return workingArray
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
