package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func getInput(filename string) ([][]byte, error) {
	var skiSlope [][]byte

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		skiSlope = append(skiSlope, []byte(reader.Text()))
	}
	return skiSlope, reader.Err()
}

func skiTheSlope(right int, down int, data [][]byte) int {
	numTrees := 0
	tree := byte(35)
	totalLength := len(data[0])
	index := 0
	for slopeLine := 0; slopeLine < len(data); slopeLine += down {
		if index >= totalLength {
			index -= totalLength
		}
		if data[slopeLine][index] == tree {
			numTrees++
		}
		index += right
	}
	return numTrees
}

func main() {
	var (
		rounds        int
		inputFilename string
	)
	flag.IntVar(&rounds, "rounds", 1, "number of rounds for benchmark")
	flag.StringVar(&inputFilename, "input", "input", "input data filename")
	flag.Parse()

	input, err := getInput(inputFilename)
	if err != nil {
		panic(err)
	}

	doTimes := make([]time.Duration, rounds)
	var res int
	for i := 0; i < rounds; i++ {
		res, doTimes[i] = doPart2(input)
	}
	var tot time.Duration
	for i := 0; i < rounds; i++ {
		tot += doTimes[i]
	}
	fmt.Printf("result:\t %v\n", res)
	fmt.Printf("elapsed:\t%v\n", tot/time.Duration(rounds))
}

func doPart2(data [][]byte) (int, time.Duration) {
	t := time.Now()
	runs := []int{1, 3, 5, 7, 1}
	rise := []int{1, 1, 1, 1, 2}
	result := 1
	for ind := 0; ind < len(runs); ind++ {
		multiplyMe := skiTheSlope(runs[ind], rise[ind], data)
		//fmt.Println(multiplyMe)
		result = result * multiplyMe
	}
	return result, time.Since(t)
}
