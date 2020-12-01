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
	for _, number1 := range numberList {
		for _, number2 := range numberList {
			for _, number3 := range numberList {
				//if number1 != number2 && number1 != number3{
				//	if number2 != number3 {
						if number1 + number2 + number3 == 2020{
							fmt.Println(number1 * number2* number3)
							fmt.Println(time.Since(t))
							os.Exit(0)
						}
					//}
				//}
			}
		}
	}

	fmt.Println(time.Since(t))
}
