package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"regexp"
	"strconv"
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
	var totalGameNumbers int
	// var enoughDice bool
	for _, line := range inputLines {
		var lineSplit []string
		var gameNumber int
		var gameCount int
		enoughDiceArray := []int{0,0,0,0,0,0,0}
		lineSplit = strings.Split(line, ": ")
		// fmt.Printf("%s\n %s\n", lineSplit[0], lineSplit[1])
		for _, findGame := range lineSplit {
			findGameRE := strings.Contains(findGame, "Game")
			check(err)
			// fmt.Printf("findGameRE: %t\n", findGameRE)
			if findGameRE == true {
				var gameNumberString string
				gameNumberString = findNumber.FindString(findGame)
				gameNumber, err = strconv.Atoi(string(gameNumberString))
				check(err)
			}
		}
		fmt.Printf("---------------Game Number: %d-------------\n", gameNumber)
		for _, games := range lineSplit {
			var game []string
			game = strings.Split(games, "; ")
			// fmt.Printf("Game Number: %d\n", gameNumber)
			gameCount := 0
			for _, colors := range game {
				var color []string
				var greenBool, redBool, blueBool bool
				var greenStr, redStr, blueStr string
				var greenDice, redDice, blueDice int
				var greenTotal, redTotal, blueTotal int
				color = strings.Split(colors, ", ")
				for _, numberColors := range color {
					// fmt.Printf("%s\n", numberColors)
					greenBool = strings.Contains(numberColors, "green")
					redBool = strings.Contains(numberColors, "red")
					blueBool = strings.Contains(numberColors, "blue")
					if greenBool == true {
						greenStr = findNumber.FindString(numberColors)
						greenDice, err = strconv.Atoi(string(greenStr))
						greenTotal = greenTotal + greenDice
						// fmt.Printf("Green Total: %d\n", greenTotal)
					}
					if redBool == true {
						redStr = findNumber.FindString(numberColors)
						redDice, err = strconv.Atoi(string(redStr))
						redTotal = redTotal + redDice
						// fmt.Printf("Red Total: %d\n", redTotal)
					}
					if blueBool == true {
						blueStr = findNumber.FindString(numberColors)
						blueDice, err = strconv.Atoi(string(blueStr))
						blueTotal = blueTotal + blueDice
						// fmt.Printf("Blue Total: %d\n", blueTotal)
					}
				}
				if greenTotal > 13 {
					// enoughDice = false
					enoughDiceArray[gameCount] = 0
					fmt.Printf("Green Failed\n")
				} else if blueTotal > 14 {
					// enoughDice = false
					fmt.Printf("Blue Failed\n")
					enoughDiceArray[gameCount] = 0
				} else if redTotal > 12 {
					// enoughDice = false
					fmt.Printf("Red Failed\n")
					enoughDiceArray[gameCount] = 0
				} else {
					// enoughDice = true
					enoughDiceArray[gameCount] = 1
				}
				fmt.Printf("enoughDiceArray:\n")
				fmt.Print(enoughDiceArray)
				gameCount = gameCount + 1
			}
		}
		var gameWorkNumber int
		gameWork := true
		for _, dice := range enoughDiceArray {
			gameWorkNumber = gameWorkNumber + dice
			gameWorkNumber = gameWorkNumber / gameCount
			fmt.Printf("gameWorkNumber: %d\n", gameWorkNumber)
			if gameWorkNumber < 1 {
				gameWork = false
			}
		}
		if gameWork == true {
			fmt.Printf("Game %d is possible!\n", gameNumber)
			totalGameNumbers = totalGameNumbers + gameNumber
		}
	}
	fmt.Print(totalGameNumbers)
}