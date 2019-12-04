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
	positionOne := (inNumber / 100000) % 10
	positionTwo := (inNumber / 10000) % 10
	positionThree := (inNumber / 1000) % 10
	positionFour := (inNumber / 100) % 10
	positionFive := (inNumber / 10) % 10
	positionSix := inNumber % 10
	if positionOne <= positionTwo && positionTwo <= positionThree && positionThree <= positionFour && positionFour <= positionFive && positionFive <= positionSix {
		intSlice := []int{positionOne, positionTwo, positionThree, positionFour, positionFive, positionSix}
		for i := 0; i <= 4; i++ {
			if intSlice[i] == intSlice[i+1] {
				return true
			}
		}
		return false
	}
	return false

}
