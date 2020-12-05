package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	planeRow = make([]int, 128)
	planeCol = []int{0, 1, 2, 3, 4, 5, 6, 7}
)

func init() {
	createPlaneRow()
}
func getInput(filename string) ([]string, error) {
	var data []string

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		data = append(data, reader.Text())
	}
	return data, reader.Err()
}

func main() {
	var (
		rounds        int
		inputFilename string
	)
	flag.IntVar(&rounds, "rounds", 1, "number of rounds for benchmark")
	flag.StringVar(&inputFilename, "input", "input.txt", "input data filename")
	flag.Parse()

	input, err := getInput(inputFilename)
	if err != nil {
		panic(err)
	}
	doTimes := make([]time.Duration, rounds)
	// PART 1
	fmt.Println("=== Part 1 ===")
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
	// PART 2
	fmt.Println("\n=== Part 2 ===")
	for i := 0; i < rounds; i++ {
		res, doTimes[i] = doPart2(input)
	}
	tot = 0
	for i := 0; i < rounds; i++ {
		tot += doTimes[i]
	}
	fmt.Printf("result:\t %v\n", res)
	fmt.Printf("elapsed:\t%v\n", tot/time.Duration(rounds))
}

func doPart1(data []string) (int, time.Duration) {
	t := time.Now()
	highestValue := 0
	for _, seatLocation := range data {
		//row, col, uniqID := findSeat(seatLocation)
		//fmt.Println(seatLocation, row, col, uniqID)
		_, _, uniqID := findSeat(seatLocation)
		//fmt.Println(seatLocation, row, col, uniqID)
		if uniqID > highestValue {
			highestValue = uniqID
		}

	}

	return highestValue, time.Since(t)
}

func doPart2(data []string) (int, time.Duration) {
	t := time.Now()
	whatamidoing := map[int][]int{}
	for x := 0; x < len(planeCol); x++ {
		whatamidoing[x] = []int{}
	}
	for _, seatLocation := range data {
		row, col, _ := findSeat(seatLocation)
		//keyValue := fmt.Sprintf("%v,%v", row, col)
		whatamidoing[row] = append(whatamidoing[row], col)
		//fmt.Println(seatLocation, row, col, uniqID)
	}
	//for ind1 := range planeRow {
	//	for ind2 := range planeCol {
	//		findKey := fmt.Sprintf("%v,%v", ind1, ind2)
	//		if _, exists := whatamidoing[findKey]; !exists {
	//			fmt.Println(findKey)
	//		}
	//	}
	//}
	for rows, cols := range whatamidoing {
		if len(cols) == 0 {
			continue
		}
		if len(cols) == 7 {
			found := false
			missingNum := 0
			for y := range cols {
				if y == cols[y] {
					found = true
				}
				if !found {
					missingNum = y
					if missingNum == 6 {
						missingNum++
					}
				}
			}
			//fmt.Println(rows, missingNum)
			return rows*8 + missingNum, time.Since(t)

		}
	}
	return 0, time.Since(t)
}

func findSeat(seatCode string) (RowNum int, ColNum int, UniqID int) {
	//F means to take the lower half, keeping rows 0 through 63.
	//B means to take the upper half,
	//var wg sync.WaitGroup
	//wg.Add(2)
	RowNum = findRow(seatCode[:7])
	ColNum = findCol(seatCode[len(seatCode)-3:])

	return RowNum, ColNum, (RowNum * 8) + ColNum
}

func createPlaneRow() {
	for i := 0; i <= 127; i++ {
		planeRow[i] = i
	}
}

func findRow(row string) int {
	workingRows := make([]int, len(planeRow))
	copy(workingRows, planeRow)
	for i := 0; i < len(row); i++ {
		if string(row[i]) == "B" {
			workingRows = workingRows[len(workingRows)/2:]
		} else {
			workingRows = workingRows[:len(workingRows)/2]
		}
	}
	return workingRows[0]
}

func findCol(col string) int {
	workingCols := make([]int, len(planeCol))
	copy(workingCols, planeCol)
	for i := 0; i < len(col); i++ {
		if string(col[i]) == "R" {
			workingCols = workingCols[len(workingCols)/2:]
		} else {
			workingCols = workingCols[:len(workingCols)/2]
		}
	}
	return workingCols[0]
}
