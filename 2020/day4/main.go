package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getInput(filename string) ([]string, error) {
	var data []string

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		data = append(data, reader.Text())
	}
	//	if len(reader.Text()) == 0 {
	//		data = append(data, subString)
	//		subString = []string{}
	//		continue
	//	}
	//	splitskis := strings.Split(reader.Text(), " ")
	//	for _, singleElement := range splitskis {
	//		subString = append(subString, singleElement)
	//	}
	//	//subString = append(subString, reader.Text())
	//}
	//data = append(data, subString)
	//fmt.Println(time.Since(t))
	return data, reader.Err()
}

func main() {
	var (
		rounds        int
		inputFilename string
	)
	flag.IntVar(&rounds, "rounds", 1, "number of rounds for benchmark")
	flag.StringVar(&inputFilename, "input", "input.txt", "input data filename")
	flag.Parse()

	input, err := getInput(inputFilename)
	//fmt.Println(input)
	if err != nil {
		panic(err)
	}
	doTimes := make([]time.Duration, rounds)
	// PART 1
	fmt.Println("=== Part 1 ===")
	var res int
	for i := 0; i < rounds; i++ {
		res, doTimes[i] = doPart1(input)
	}
	var tot time.Duration
	for i := 0; i < rounds; i++ {
		tot += doTimes[i]
	}
	fmt.Printf("result:\t %v\n", res)
	fmt.Printf("elapsed:\t%v\n", tot/time.Duration(rounds))
	// PART 2
	//fmt.Println("\n=== Part 2 ===")
	//for i := 0; i < rounds; i++ {
	//	res, doTimes[i] = doPart2(input)
	//}
	//tot = 0
	//for i := 0; i < rounds; i++ {
	//	tot += doTimes[i]
	//}
	//fmt.Printf("result:\t %v\n", res)
	//fmt.Printf("elapsed:\t%v\n", tot/time.Duration(rounds))
}

func parseInput(input [][]string) map[int]map[string]string {
	parsedData := make(map[int]map[string]string)
	for ind, data := range input {
		parsedData[ind] = make(map[string]string)

		for _, values := range data {
			for _, ppInfo := range strings.Split(values, " ") {
				keyValue := strings.Split(ppInfo, ":")
				parsedData[ind][keyValue[0]] = keyValue[1]
			}
		}

	}

	return parsedData
}

func doPart1(data []string) (int, time.Duration) {
	t := time.Now()
	//organizedData := parseInput(data)
	valid := 0
	for _, line := range data {
		if len(line) == 0 {

		}
		fmt.Println(len(line))
	}
	//for _, passports := range organizedData {
	//	if len(passports) == 8 {
	//		valid++
	//	} else if len(passports) == 7 {
	//		if _, exists := passports["cid"]; !exists {
	//			valid++
	//		}
	//	}
	//}
	//for i := 0; i < len(data); i++ {
	//	if len(data[i]) == 8 {
	//		valid++
	//		continue
	//	} else {
	//		if len(data[i]) == 7 {
	//			exists := 0
	//			for _, entry := range data[i] {
	//				if strings.HasPrefix(entry, "cid:") {
	//					exists++
	//				}
	//			}
	//			if exists == 0 {
	//				valid++
	//			}
	//		}
	//	}
	//}
	//	count := 0
	//	for _, values := range strings.Split(stuff, ":") {
	//		switch keyValue[0] {
	//		case "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid":
	//			count++
	//		case "cid":
	//			count--
	//		}
	//		count := 0
	//		fmt.Println(values)
	//		fmt.Println(len(strings.Split(values, " ")))
	//		for _, ppInfo := range strings.Split(values, " ") {
	//			//fmt.Println(ppInfo)
	//			keyValue := strings.Split(ppInfo, ":")
	//			switch keyValue[0] {
	//			case "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid":
	//				count++
	//			case "cid":
	//				count--
	//			}
	//		}
	//		//fmt.Println(count)
	//		if count == 7 {
	//			valid++
	//		}
	//	}
	//}
	return valid, time.Since(t)
}

func doPart2(data [][]string) (int, time.Duration) {
	t := time.Now()
	organizedData := parseInput(data)
	valid := 0
	for _, passports := range organizedData {
		if validateData(passports) {
			valid++
		}
	}
	return valid, time.Since(t)
}

func validateData(checkMe map[string]string) bool {
	_, exists := checkMe["cid"]

	if len(checkMe) == 8 || (len(checkMe) == 7 && !exists) {
		//byr (Birth Year) - four digits; at least 1920 and at most 2002.
		//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
		//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
		//hgt (Height) - a number followed by either cm or in:
		//If cm, the number must be at least 150 and at most 193.
		//If in, the number must be at least 59 and at most 76.
		//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
		//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
		//pid (Passport ID) - a nine-digit number, including leading zeroes.
		//fmt.Println(checkMe)
		for key, value := range checkMe {
			switch key {
			case "byr":
				//fmt.Println(value)
				if len(value) != 4 {
					//fmt.Println("Failed byr")
					return false
				}
				date, err := strconv.Atoi(value)
				if err != nil {
					//fmt.Println("Failed byr")
					return false
				}
				if date < 1920 || date > 2002 {
					//fmt.Println("whaa")
					//fmt.Println("Failed byr")
					return false
				}
			case "iyr":
				if len(value) != 4 {
					//fmt.Println("Failed iyr")
					return false
				}
				date, err := strconv.Atoi(value)
				if err != nil {
					//fmt.Println("Failed iyr")
					return false
				}
				if date < 2010 || date > 2020 {
					//fmt.Println("Failed iyr")
					return false
				}
			case "eyr":
				if len(value) != 4 {
					//fmt.Println("Failed eyr")
					return false
				}
				date, err := strconv.Atoi(value)
				if err != nil {
					//fmt.Println("Failed eyr")
					return false
				}
				if date < 2020 || date > 2030 {
					return false
				}
			case "hgt":
				//fmt.Println(value[:len(value)-2])
				height, err := strconv.Atoi(value[:len(value)-2])
				if err != nil {
					//fmt.Println("Failed hgt")
					return false
				}
				unit := value[len(value)-2:]
				var lowRange, highRange = 150, 193
				if unit == "in" {
					lowRange = 59
					highRange = 76
				}
				if height < lowRange || height > highRange {
					return false
				}
			case "hcl":
				//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
				if !strings.HasPrefix(value, "#") || len(value) != 7 {
					return false
				}
				for i := 1; i < len(value); i++ {
					switch string(value[i]) {
					case "a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
						continue
					default:
						return false
					}

				}
			case "ecl":
				switch value {
				case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
					continue
				default:
					//fmt.Println("Failed ecl")
					return false
				}
			case "pid":
				if len(value) != 9 {
					//fmt.Println("Failed pid")
					return false
				}
				_, err := strconv.Atoi(value)
				if err != nil {
					//fmt.Println("Failed pid")
					return false
				}
			}
		}
		return true
	}
	//fmt.Println(checkMe)
	return false
}
