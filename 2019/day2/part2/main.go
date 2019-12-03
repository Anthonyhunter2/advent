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
	//t := time.Now()
	inputFile := os.Args[1]
	numberList := []int{}
	listNums := readInFile(inputFile)
	for _, numString := range listNums {
		num, err := strconv.Atoi(numString)
		if err != nil {
			fmt.Println(err)
		}
		numberList = append(numberList, num)
	}
	inputInts = numberList
	t := time.Now()
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			workingInput := make([]int, len(inputInts))
			copy(workingInput, inputInts)
			workingInput[1] = x
			workingInput[2] = y
			Updatelist(0, workingInput)
			if workingInput[0] == 19690720 {
				fmt.Println(workingInput[1], " ", workingInput[2])
				fmt.Println(100*workingInput[1] + workingInput[2])
				fmt.Println(time.Since(t))
				break
			}
		}
	}
	fmt.Printf("Total Execution time: %v\n", time.Since(t))
}

func Updatelist(ind int, workingInput []int) {
	value1 := workingInput[ind+1]
	value2 := workingInput[ind+2]
	value3 := workingInput[ind+3]
	if workingInput[ind] == 99 {
		return
	} else if workingInput[ind] == 1 {
		workingInput[value3] = workingInput[value1] + workingInput[value2]
	} else if inputInts[ind] == 2 {
		workingInput[value3] = workingInput[value1] * workingInput[value2]
	} else {
		fmt.Println("You messed up")
		os.Exit(1)
	}
	Updatelist(ind+4, workingInput)
}
func readInFile(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		listNums := strings.Split(reader.Text(), ",")
		return listNums
	}
	return []string{}
}
