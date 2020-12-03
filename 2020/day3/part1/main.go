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

func main() {
	data := getInput()
	t := time.Now()
	tree := byte(35)
	numTrees := 0
	// The -1 keeps our index values inline with the length of our data
	totalLength := len(data[0]) - 1
	index := 0

	for slopeLine := range data {
		if slopeLine == 0 {
			continue
		}
		if totalLength-index < 3 {
			index = 3 - (totalLength - index) - 1
		} else {
			index += 3
		}
		if data[slopeLine][index] == tree {
			numTrees++
		}
	}

	fmt.Println(numTrees)
	fmt.Println(time.Since(t))
}
