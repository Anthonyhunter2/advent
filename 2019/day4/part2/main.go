package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	count := 0
	puzzleInputMin := 146810
	puzzleInputMax := 612564
	//possiblePasswords := []int
	for i := puzzleInputMin + 1; i < puzzleInputMax; i++ {
		if checkDecrease(i) {
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(time.Since(t))
}

func checkDecrease(inNumber int) bool {
	pZero := (inNumber / 100000) % 10
	pOne := (inNumber / 10000) % 10
	pTwo := (inNumber / 1000) % 10
	pThree := (inNumber / 100) % 10
	pFour := (inNumber / 10) % 10
	pFive := inNumber % 10
	if pZero <= pOne && pOne <= pTwo && pTwo <= pThree && pThree <= pFour && pFour <= pFive {
		intSlice := []int{pZero, pOne, pTwo, pThree, pFour, pFive}
		counter := make(map[int]int)
		for _, value := range intSlice {
			counter[value]++
		}
		for _, v := range counter {
			if v == 2 {
				return true
			}
		}

		return false
	}
	return false
}
