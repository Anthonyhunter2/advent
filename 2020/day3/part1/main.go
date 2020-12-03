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
		res, doTimes[i] = doPart1(input)
	}
	var tot time.Duration
	for i := 0; i < rounds; i++ {
		tot += doTimes[i]
	}
	fmt.Printf("result:\t %v\n", res)
	fmt.Printf("elapsed:\t%v\n", tot/time.Duration(rounds))
}

func doPart1(data [][]byte) (int, time.Duration) {
	t := time.Now()
	tree := byte(35)
	numTrees := 0
	// The -1 keeps our index values inline with the length of our data
	totalLength := len(data[0])
	index := 0

	for slopeLine := range data {
		fmt.Println(data[slopeLine], index)
		if slopeLine == 0 {
			index += 3
			continue
		}
		if index > totalLength {
			index -= totalLength
		}
		if data[slopeLine][index] == tree {
			numTrees++
		}

		index += 3
	}
	return numTrees, time.Since(t)
}
