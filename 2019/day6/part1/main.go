package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var directOribits = map[string][]string{}
var count = 0

func main() {

	t := time.Now()
	inputFile := os.Args[1]
	readInFile(inputFile)
	//workingInput := make([]int, len(inputInts))
	//copy(workingInput, inputInts)
	startingKey := findStart()
	CountALLOrbits(startingKey)
	fmt.Println("Count ", count)
	fmt.Println(directOribits)
	fmt.Println(time.Since(t))
}

func readInFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ")")
		if _, found := directOribits[line[0]]; !found {
			directOribits[line[0]] = append(directOribits[line[0]], line[1])
		} else if found {
			directOribits[line[0]] = append(directOribits[line[0]], line[1])
		} else if _, ok := directOribits[line[1]]; !ok {
			directOribits[line[1]] = []string{}
		}
	}
}

func findStart() (startingPoint string) {
	count := 0
	initalKeys := map[string]struct{}{}

	for _, v := range directOribits {
		for i := 0; i < len(v); i++ {
			count++
			//if _, found := directOribits[v[i]]; !found {
			//	fmt.Println(v[i])
			//}
			initalKeys[v[i]] = struct{}{}
		}
	}

	for k, _ := range directOribits {
		if _, found := initalKeys[k]; !found {
			return k
		}
	}
	return ""
}

func CountALLOrbits(startingKey string) {
	startingList := directOribits[startingKey]
	for i := 0; i < len(startingList); i++ {
		count++
		fmt.Println(startingList[i])
		CountALLOrbits(startingList[i])
	}
}
