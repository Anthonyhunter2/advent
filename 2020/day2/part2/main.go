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
	}
	return pwList
}
func main() {
	//t := time.Now()
	data := getInput()
	t := time.Now()
	goodPWs := 0
	for _, password := range data {
		splitLine := strings.Split(password, " ")
		keyLetter := strings.TrimRight(splitLine[1], ":")
		ranges := strings.Split(splitLine[0], "-")
		rangeLow, err := strconv.Atoi(ranges[0])
		if err != nil {
			fmt.Println(err)
		}
		rangeHigh, err := strconv.Atoi(ranges[1])
		if err != nil {
			fmt.Println(err)
		}
		ind1 := string(splitLine[2][rangeLow-1])
		ind2 := string(splitLine[2][rangeHigh-1])
		if (ind1 != keyLetter) != (ind2 != keyLetter) {
			goodPWs++
		}
		//switch {
		//case ind1 == keyLetter && ind2 == keyLetter:
		//	continue
		//case ind1 == keyLetter:
		//	goodPWs += 1
		//case ind2 == keyLetter:
		//	goodPWs += 1
		//}
	}
	fmt.Println(goodPWs)
	fmt.Println(time.Since(t))
}
