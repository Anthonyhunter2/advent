package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	numberList := []float64{}
	total := 0.0
	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		stringNum, err := strconv.ParseFloat(reader.Text(), 64)
		if err != nil {
			fmt.Println(err)
		}
		numberList = append(numberList, stringNum)
	}
	for _, number := range numberList {
		addMe := calcFuel(number)
		for addMe > 0 {
			total = total + addMe
			addMe = calcFuel(addMe)
		}
	}
	fmt.Println(total)
}

func calcFuel(incoming float64) float64 {
	checkThis := math.Floor(incoming/3) - 2
	fmt.Println(checkThis)
	return checkThis

}
