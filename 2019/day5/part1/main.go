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
var count = 0

func main() {
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
	count++
	value1 := inputInts[ind+1]
	value2 := inputInts[ind+2]
	value3 := inputInts[ind+3]

	if inputInts[ind] < 10 {
		mainLoop(ind, value1, value2, value3)
	} else {
		firstValue := 0
		secondValue := 0
		positionMode2 := 0
		if inputInts[ind] > 1000 {
			positionMode2 = (inputInts[ind] / 1000) % 10
		}
		positionMode1 := (inputInts[ind] / 100) % 10
		shouldBeZero := (inputInts[ind] / 10) % 10
		action := inputInts[ind] % 10

		if shouldBeZero != 0 {
			fmt.Printf("I guess i dont understand, %d", inputInts[ind])
			os.Exit(1)
		}
		if positionMode1 == 0 {
			firstValue = possitionMode(value1)
		} else if positionMode1 == 1 {
			firstValue = immediateMode(value1)
		}
		if action == 1 {
			if positionMode2 == 0 {
				secondValue = possitionMode(value2)
			} else if positionMode2 == 1 {
				secondValue = immediateMode(value2)
			}
		} else if action == 2 {
			if positionMode2 == 0 {
				secondValue = possitionMode(value2)
			} else if positionMode2 == 1 {
				secondValue = immediateMode(value2)
			}
		}
		intcode(ind, firstValue, secondValue, value3, action)
	}
}

func mainLoop(ind int, value1 int, value2 int, value3 int) {

	action := inputInts[ind]
	if action == 99 {
		return
	} else if action == 4 {
		printValAt := value1
		if inputInts[ind+2] == 99 {
			fmt.Println("Diagnostic code", inputInts[printValAt])
			return
		}
		if inputInts[printValAt] != 0 {
			fmt.Println("You messed up on step", count)
			fmt.Println(inputInts[ind-4], inputInts[ind-3], inputInts[ind-2], inputInts[ind-1], printValAt)
			os.Exit(1)
		}
		Updatelist(ind + 2)

	} else if action == 3 && ind == 0 {
		replaceNumat := inputInts[ind+1]
		inputInts[replaceNumat] = 1
		Updatelist(ind + 2)
	} else if action == 1 {
		inputInts[value3] = inputInts[value1] + inputInts[value2]
		Updatelist(ind + 4)
	} else if action == 2 {
		inputInts[value3] = inputInts[value1] * inputInts[value2]
		Updatelist(ind + 4)
	}

}
func intcode(ind int, value1 int, value2 int, value3 int, action int) {
	if action == 99 {
		return
	} else if action == 4 {
		printValAt := value1
		if inputInts[ind+2] == 99 {
			fmt.Println("Diagnostic code", inputInts[printValAt])
			return
		}
		if printValAt != 0 {
			fmt.Println("You messed up on step", count-1, "validationCode:", printValAt)
			fmt.Println(inputInts[225])
			fmt.Println(inputInts[ind-4], inputInts[ind-3], inputInts[ind-2], inputInts[ind-1], printValAt)
			os.Exit(1)
		}
		Updatelist(ind + 2)

	} else if action == 3 && ind == 0 {
		replaceNumat := inputInts[ind+1]
		inputInts[replaceNumat] = 1
		Updatelist(ind + 2)
	} else if action == 1 {
		inputInts[value3] = value1 + value2
		Updatelist(ind + 4)
	} else if action == 2 {
		inputInts[value3] = value1 * value2
		Updatelist(ind + 4)
	}
}
func immediateMode(ind int) int {
	return ind
}
func possitionMode(ind int) int {
	return inputInts[ind]
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
