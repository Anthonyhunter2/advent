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
	Updatelist(0)
	fmt.Printf("Total Execution time: %v\n", time.Since(t))
}

func Updatelist(ind int) {
	value1 := inputInts[ind+1]
	value2 := inputInts[ind+2]
	value3 := inputInts[ind+3]
	if inputInts[ind] == 99 {
		fmt.Println(inputInts[0])
		return
	} else if inputInts[ind] == 1 {
		inputInts[value3] = inputInts[value1] + inputInts[value2]
	} else if inputInts[ind] == 2 {
		inputInts[value3] = inputInts[value1] * inputInts[value2]
	} else {
		fmt.Println("You messed up")
		os.Exit(1)
	}
	Updatelist(ind + 4)
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
