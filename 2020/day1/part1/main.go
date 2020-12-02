package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var numberList []int

	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		stringNum, err := strconv.Atoi(reader.Text())
		if err != nil {
			fmt.Println(err)
		}
		numberList = append(numberList, stringNum)
	}
	t := time.Now()
	//for _, num1 := range numberList {
	//	for _, num2 := range numberList {
	//		if num1+num2 == 2020 {
	//			fmt.Println(num1 * num2)
	//			fmt.Println(time.Since(t))
	//			os.Exit(0)
	//		}
	//	}
	//}

	for ind := 0; ind < len(numberList); ind++ {
		for _, num := range numberList[ind:] {
			if numberList[ind]+num == 2020 {
				fmt.Println(numberList[ind] * num)
				fmt.Println(time.Since(t))
				return
			}
		}
	}
}
