package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var numberList []float64

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
	t := time.Now()
	for _, number1 := range numberList {
		for _, number2 := range numberList { {
		if number1 != number2{
		if number1 + number2 == 2020{
			fmt.Println(number1 * number2)
			fmt.Println(time.Since(t))
			os.Exit(0)
		}}
		}
		}
	}

	fmt.Println(time.Since(t))
}
