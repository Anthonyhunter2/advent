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
var startingPoint = 5
var t time.Time

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
	t = time.Now()
	Updatelist(0)
	fmt.Printf("Total Execution time: %v\n", time.Since(t))
}

func Updatelist(ind int) {
	count++
	//fmt.Println("main Count", count)
	//fmt.Println("before work:", inputInts)
	positionMode1 := 0
	positionMode2 := 0
	firstValue := 0
	secondValue := 0
	value1 := inputInts[ind+1]
	value2 := 0
	value3 := 0

	if inputInts[ind] < 10 {
		if inputInts[ind] != 4 && inputInts[ind] != 3 {
			value3 = inputInts[ind+3]
			value2 = inputInts[ind+2]
		}
		//fmt.Println(value1, "<<<", value2)
		firstValue = possitionMode(value1)
		secondValue = possitionMode(value2)
		mainLoop(ind, firstValue, secondValue, value3, 0)
		fmt.Println(time.Since(t))
		os.Exit(0)
		goto END
		return
	} else {
		if inputInts[ind] > 1000 {
			positionMode2 = (inputInts[ind] / 1000) % 10
		}
		if inputInts[ind] != 4 && inputInts[ind] != 3 {
			value3 = inputInts[ind+3]
			value2 = inputInts[ind+2]
		}
		positionMode1 = (inputInts[ind] / 100) % 10
		shouldBeZero := (inputInts[ind] / 10) % 10
		action := inputInts[ind] % 10

		if shouldBeZero != 0 {
			fmt.Printf("I guess i dont understand, %d\n%v\n", inputInts[ind], count)
			fmt.Println(inputInts[ind-4], inputInts[ind-3], inputInts[ind-2], inputInts[ind-1])
			os.Exit(1)
		}
		if positionMode1 == 0 {
			firstValue = possitionMode(value1)
		} else if positionMode1 == 1 {
			firstValue = immediateMode(value1)
		}
		if action != 4 {
			if positionMode2 == 0 {
				secondValue = possitionMode(value2)
			} else if positionMode2 == 1 {
				secondValue = immediateMode(value2)
			}
		}
		mainLoop(ind, firstValue, secondValue, value3, action)
	}
END:
	return
}

func mainLoop(ind int, value1 int, value2 int, value3 int, inComingAction int) {
	action := 0
	if inputInts[ind] > 10 {
		action = inComingAction
	} else {
		action = inputInts[ind]
	}
	if action == 99 {
		return
	} else if action == 4 {
		printValAt := value1
		if inputInts[ind+2] == 99 {
			//fmt.Println(inputInts)
			fmt.Println("Diagnostic code", printValAt)
			//os.Exit(0)
			return
		}
		if printValAt != 0 {
			fmt.Println("You messed up on step", count)
			fmt.Println("diag:", printValAt)
			fmt.Println(inputInts[ind-4], inputInts[ind-3], inputInts[ind-2], inputInts[ind-1], printValAt)
			os.Exit(1)
		}
		Updatelist(ind + 2)

	} else if action == 3 && ind == 0 {
		replaceNumat := inputInts[ind+1]
		inputInts[replaceNumat] = startingPoint
		Updatelist(ind + 2)
	} else if action == 1 {
		inputInts[value3] = value1 + value2
		Updatelist(ind + 4)
	} else if action == 2 {
		inputInts[value3] = value1 * value2
		Updatelist(ind + 4)
	} else if action == 5 {
		if value1 != 0 {
			Updatelist(value2)
		}
		Updatelist(ind + 3)
	} else if action == 6 {
		if value1 == 0 {
			Updatelist(value2)
		}
		Updatelist(ind + 3)
	} else if action == 7 {
		if value1 < value2 {
			inputInts[value3] = 1
			Updatelist(ind + 4)
		}
		inputInts[value3] = 0
		Updatelist(ind + 4)
	} else if action == 8 {
		//fmt.Println(value1, ",", value2)
		if value1 == value2 {
			inputInts[value3] = 1
			Updatelist(ind + 4)
		}
		inputInts[value3] = 0
		Updatelist(ind + 4)
	}
	return

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
