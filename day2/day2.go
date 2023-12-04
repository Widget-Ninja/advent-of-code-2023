package main

import (
	"fmt"
	"os"
	"bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	input, err := os.Open("C:/Users/shamu/Documents/Code/advent-of-code-2023/day2/input.txt")
	check(err)

	fileScan := bufio.NewScanner(input)
    fileScan.Split(bufio.ScanLines)
	var inputLines []string
    for fileScan.Scan() {
        inputLines = append(inputLines, fileScan.Text())
    }
	fmt.Print(inputLines)
}