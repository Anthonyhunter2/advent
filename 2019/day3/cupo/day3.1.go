package main
import (
	"time"
	"math"
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
func main(){
	input := os.Args[1]
	defaultCode := get_code(input)
	t1 := time.Now()
	m := make(map[int][]int)
	run_code(defaultCode[0], m)
	mdist := check_dist(defaultCode[1], m)
	fmt.Println("Manhattan Distance:",mdist)
	fmt.Println("Time:",time.Since(t1))
}
func get_code(filename string) [][]string{
	file,err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()
	var tmp[][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		tmp = append(tmp, strings.Split(scanner.Text(), ","))
	}
	return tmp
}
func run_code(code []string, m map[int][]int){
	x,y := 0,0
	if _, exists := m[0]; !exists{
		m[0] = []int{0}
	}
	for commandNum:=0; commandNum<len(code); commandNum++{
		command := []rune(code[commandNum])
		direction := string(command[0])
		distance,_ := strconv.Atoi(string(command[1:]))

		trace_line(&x, &y, direction, distance, m)
	}
}
func trace_line(x *int, y *int, dir string, dist int, m map[int][]int){
	for i:=dist; i>0; i--{
		switch dir{
		case "U":
			*y++
		case "D":
			*y--
		case "R":
			*x++
		case "L":
			*x--
		default:
			fmt.Println("No matching case for direction:",dir)
			os.Exit(1)
		}
		draw(*x, *y, m)
	}
}
func draw(x int, y int, m map[int][]int){
	m[x] = append(m[x],y)
}
func check(x int, y int, m map[int][]int) bool {
	for _, item := range m[x]{
		if item == y{
			return true
		}
	}
	return false
}
func man_dist(x1 int, y1 int, x2 int, y2 int) int{
	return int(math.Abs(float64(x1-x2))+math.Abs(float64(y1-y2)))
}
func check_dist(code []string, m map[int][]int) int{
	x,y,minDist := 0,0,0
        for commandNum:=0; commandNum<len(code); commandNum++{
                command := []rune(code[commandNum])
                direction := string(command[0])
                distance,_ := strconv.Atoi(string(command[1:]))

		for ;distance > 0; distance--{
			switch direction{
			case "U":
				y++
			case "D":
				y--
			case "R":
				x++
			case "L":
				x--
			default:
				fmt.Println("No matching case for direction:",direction)
				os.Exit(1)
			}
			if check(x,y,m){
				tmp := man_dist(0, 0, x, y)
				if tmp < minDist{
					minDist = tmp
				}else if minDist == 0{
					minDist = tmp
				}
			}
		}
	}
	return minDist
}
