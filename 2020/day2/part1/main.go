package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type PWEntry struct {
	RangeLow  int
	RangeHigh int
	KeyLetter string
	Password  string
}

func getInput() []PWEntry {
	var pwList []PWEntry

	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		newEntry := PWEntry{}
		splitLine := strings.Split(reader.Text(), ":")
		keyValues := strings.Split(splitLine[0], " ")
		rangeLow, err := strconv.Atoi(strings.Split(keyValues[0], "-")[0])
		if err != nil {
			fmt.Println(err)
		}
		rangeHigh, err := strconv.Atoi(strings.Split(keyValues[0], "-")[1])
		if err != nil {
			fmt.Println(err)
		}
		newEntry.RangeLow = rangeLow
		newEntry.RangeHigh = rangeHigh
		newEntry.KeyLetter = keyValues[1]
		newEntry.Password = splitLine[1]
		pwList = append(pwList, newEntry)
	}
	return pwList
}
func main() {
	t := time.Now()
	data := getInput()
	goodPWs := 0
	for _, password := range data {
		numberOfInstancesOfKey := strings.Count(password.Password, password.KeyLetter)
		if password.RangeLow <= numberOfInstancesOfKey && numberOfInstancesOfKey <= password.RangeHigh {
			goodPWs += 1
		}
	}
	fmt.Println(goodPWs)
	fmt.Println(time.Since(t))
}
