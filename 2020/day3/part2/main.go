package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func getInput() [][]byte {
	var skiSlope [][]byte

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		skiSlope = append(skiSlope, []byte(reader.Text()))
	}
	return skiSlope
}

func skiTheSlope(right int, down int, data [][]byte) int {
	numTrees := 0
	tree := byte(35)
	totalLength := len(data[0]) - 1
	index := 0
	slopeLine := 0
	for range data {
		if slopeLine == 0 {
			slopeLine += down
			continue
		} else if slopeLine > len(data) {
			return numTrees
		}
		if totalLength-index < right {
			index = right - (totalLength - index) - 1
		} else {
			index += right
		}
		if data[slopeLine][index] == tree {
			numTrees++
		}
		slopeLine += down
	}
	return numTrees
}

func main() {
	data := getInput()
	t := time.Now()

	runs := []int{1, 3, 5, 7, 1}
	rise := []int{1, 1, 1, 1, 2}
	result := 1
	for ind := 0; ind < len(runs); ind++ {
		multiplyMe := skiTheSlope(runs[ind], rise[ind], data)
		//fmt.Println(multiplyMe)
		result = result * multiplyMe
	}

	fmt.Println(result)
	fmt.Println(time.Since(t))
}
