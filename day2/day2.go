package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"regexp"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	input, err := os.Open("C:/Users/shamu/Documents/Code/advent-of-code-2023/day2/inputTest.txt")
	check(err)
	findNumber := regexp.MustCompile("[0-9]+")

	fileScan := bufio.NewScanner(input)
    fileScan.Split(bufio.ScanLines)
	var inputLines []string
    for fileScan.Scan() {
        inputLines = append(inputLines, fileScan.Text())
    }
	for _, line := range inputLines {
		var lineSplit []string
		lineSplit = strings.Split(line, ": ")
		// fmt.Printf("%s\n %s\n", lineSplit[0], lineSplit[1])
		
		for _, games := range lineSplit {
			var game []string
			game = strings.Split(games, "; ")
			// for _, numberGames := range game {
			// fmt.Printf("%s\n", numberGames)
			// }
			for _, colors := range game {
				var color []string
				color = strings.Split(colors, ", ")
				for _, numberColors := range color {
					fmt.Printf("%s\n", numberColors)
				}
				findNumber := regexp.MustCompile("[0-9]+")
				// findGreen := regexp.MustCompile(`green`)
				// findRed := regexp.MustCompile(`red`)
				// findBlue := regexp.MustCompile(`blue`)
			}
		}
	}
	// fmt.Print(inputLines)
}