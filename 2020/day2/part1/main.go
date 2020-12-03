package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getInput() []string {
	var pwList []string

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		pwList = append(pwList, reader.Text())
		//newEntry := PWEntry{}
		//splitLine := strings.Split(reader.Text(), ":")
		//keyValues := strings.Split(splitLine[0], " ")
		//rangeLow, err := strconv.Atoi(strings.Split(keyValues[0], "-")[0])
		//if err != nil {
		//	fmt.Println(err)
		//}
		//rangeHigh, err := strconv.Atoi(strings.Split(keyValues[0], "-")[1])
		//if err != nil {
		//	fmt.Println(err)
		//}
		//newEntry.RangeLow = rangeLow
		//newEntry.RangeHigh = rangeHigh
		//newEntry.KeyLetter = keyValues[1]
		//newEntry.Password = splitLine[1]
		//pwList = append(pwList, newEntry)
	}
	return pwList
}
func main() {
	data := getInput()
	t := time.Now()

	goodPWs := 0
	for _, password := range data {
		splitLine := strings.Split(password, ":")
		keyValues := strings.Split(splitLine[0], " ")
		ranges := strings.Split(keyValues[0], "-")
		rangeLow, err := strconv.Atoi(ranges[0])
		if err != nil {
			fmt.Println(err)
		}
		rangeHigh, err := strconv.Atoi(ranges[1])
		if err != nil {
			fmt.Println(err)
		}
		numberOfInstancesOfKey := strings.Count(splitLine[1], keyValues[1])
		if rangeLow <= numberOfInstancesOfKey && numberOfInstancesOfKey <= rangeHigh {
			goodPWs += 1
		}
	}
	fmt.Println(goodPWs)
	fmt.Println(time.Since(t))
}
