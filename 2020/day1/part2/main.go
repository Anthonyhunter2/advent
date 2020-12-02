package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func getInput() []int {
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
	return numberList
}
func main() {

	numberList := getInput()
	t := time.Now()

	for _, num1 := range numberList {
		for _, num2 := range numberList {
			for _, num3 := range numberList {
				if num1+num2+num3 == 2020 {
					et := time.Now().Sub(t)
					fmt.Println(num1 * num2 * num3)
					fmt.Println(et)
					return
				}
			}
		}
	}

}
