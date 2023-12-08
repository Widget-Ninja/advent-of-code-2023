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
	input, err := os.Open("C:/Users/shamu/Documents/Code/advent-of-code-2023/day2/input.txt")
	check(err)
	
	findNumber := regexp.MustCompile("[0-9]+")
	fileScan := bufio.NewScanner(input)
    fileScan.Split(bufio.ScanLines)
	var inputLines []string
    for fileScan.Scan() {
        inputLines = append(inputLines, fileScan.Text())
    }
	// var totalGameNumbers int
	var gameWorkNumber int
	// var enoughDice bool
	for _, line := range inputLines {
		var lineSplit []string
		var gameNumber int
		var gameCount int
		var powerDice int
		// minDiceArray := []int{0, 0, 0, 0}
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
			gameCount = 0
			greenTotal, redTotal, blueTotal := 0, 0, 0
			for _, colors := range game {
				var color []string
				var greenBool, redBool, blueBool bool
				var greenStr, redStr, blueStr string
				var greenDice, redDice, blueDice int
				
				color = strings.Split(colors, ", ")
				for _, numberColors := range color {
					// fmt.Printf("%s\n", numberColors)
					greenBool = strings.Contains(numberColors, "green")
					redBool = strings.Contains(numberColors, "red")
					blueBool = strings.Contains(numberColors, "blue")
					if greenBool == true {
						greenStr = findNumber.FindString(numberColors)
						greenDice, err = strconv.Atoi(string(greenStr))
						// greenTotal = greenTotal + greenDice
						// fmt.Printf("Green Total: %d\n", greenTotal)
					}
					if redBool == true {
						redStr = findNumber.FindString(numberColors)
						redDice, err = strconv.Atoi(string(redStr))
						// redTotal = redTotal + redDice
						// fmt.Printf("Red Total: %d\n", redTotal)
					}
					if blueBool == true {
						blueStr = findNumber.FindString(numberColors)
						blueDice, err = strconv.Atoi(string(blueStr))
						// blueTotal = blueTotal + blueDice
						// fmt.Printf("Blue Total: %d\n", blueTotal)
					}
				}
				if greenDice > greenTotal {
					// enoughDice = false
					greenTotal = greenDice
					fmt.Printf("Green Updated to %d\n", greenTotal)
				}
				if blueDice > blueTotal {
					// enoughDice = false
					blueTotal = blueDice
					fmt.Printf("Blue Updated to %d\n", blueTotal)
				}
				if redDice > redTotal {
					// enoughDice = false
					redTotal = redDice
					fmt.Printf("Red Updated to %d\n", greenTotal)
				} 
				// fmt.Printf("minDiceArray:\n")
				gameCount = gameCount + 1
			}
			fmt.Printf("green: %d, blue: %d, red: %d\n", greenTotal, blueTotal, redTotal)
			powerDice = greenTotal * blueTotal * redTotal
			fmt.Printf("powerDice: %d\n", powerDice)
			// arrayAppend := []int{0}
			// minDiceArray = append(minDiceArray, arrayAppend...)
			// minDiceArray[gameNumber] = powerDice
			// fmt.Printf("minDiceArray: %v\n", minDiceArray)
		}
		
		// gameWork := true
		// for _, dice := range minDiceArray {
			// fmt.Printf("dice: %d\n", dice)
		gameWorkNumber = gameWorkNumber + powerDice
			// fmt.Printf("Before division gameWorkNumber: %d\n", gameWorkNumber)
		// }
		fmt.Printf("gameCount: %d\n", gameCount)
		// gameWorkNumber = gameWorkNumber / gameCount
		fmt.Printf("gameWorkNumber: %d\n", gameWorkNumber)
		// if gameWorkNumber == 1 {
		// 	gameWork = true
		// } else {
		// 	gameWork = false
		// }
		// fmt.Printf("gameWork: %t\n", gameWork)
		// if gameWork == true {
		// 	fmt.Printf("Game %d is possible!\n", gameNumber)
		// 	totalGameNumbers = totalGameNumbers + gameNumber
		// }
	}
	fmt.Print(gameWorkNumber)
}