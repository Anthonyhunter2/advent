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
		total = total + math.Floor(number/3) - 2
	}
	fmt.Println(total)
}
