package main

import (
	"fmt"
	"os"
	// "regexp"
	"bufio"
	// "strings"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main () {
	input, err := os.Open("C:/Users/shamu/Documents/Code/advent-of-code-2023/day1/input.txt")
	check(err)
	// re := regexp.MustCompile("[0-9]+")
	// fmt.Println(re.FindAllString(string(input), -1))
    fileScan := bufio.NewScanner(input)
    fileScan.Split(bufio.ScanLines)
    var coordinatesLines []string
	// var numbers

    for fileScan.Scan() {
		// numbers = re.FindAllString(string(fileScan.Text()), -1)
        coordinatesLines = append(coordinatesLines, fileScan.Text())
    }
	// numbers := len(strings.Fields(coordinatesLines))
	fmt.Print(len(coordinatesLines))
	// for coordinatesLines {

	// }
}