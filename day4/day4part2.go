package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	input, err := os.Open("C:/Users/shamu/Documents/Code/advent-of-code-2023/day4/input.txt")
	check(err)

	fileScan := bufio.NewScanner(input)
    fileScan.Split(bufio.ScanLines)
	var inputLines []string
    for fileScan.Scan() {
        inputLines = append(inputLines, fileScan.Text())
    }
	// fmt.Print(inputLines)
	totalScore := 0
	cardCount := 1
	for _, line := range inputLines {
		var lineSplit []string
		lineSplit = strings.Split(line, ": ")
		// fmt.Printf("%s\n", lineSplit)
		lineSplit = strings.Split(lineSplit[1], " | ")
		// fmt.Printf("%s\n", lineSplit)
		winner := lineSplit[0]
		// fmt.Printf("%s\n", winner)
		winnerNumbers := strings.Split(winner, " ")
		// fmt.Printf("%s\n", winnerNumbers[0])
		card := lineSplit[1]
		// fmt.Printf("%s\n", card)
		cardNumbers := strings.Split(card, " ")
		// fmt.Printf("%s\n", cardNumbers[0])
		cardScore := 0
		for _, cardNumber := range cardNumbers {
			// fmt.Printf("'%s'\n", cardNumber)
			var intCard int
			if cardNumber != "" {
				// fmt.Printf("card number '%s'\n", cardNumber)
				intCard, err = strconv.Atoi(cardNumber)
				check(err)
			}
			for _, winnerNumber := range winnerNumbers {
				var intWinner int
				if winnerNumber != "" {
					// fmt.Printf("Test '%s', against '%s'\n", cardNumber, winnerNumber)
					intWinner, err = strconv.Atoi(winnerNumber)
					check(err)
				}
				if intCard != 0 {
					if intWinner != 0 {
						if intCard == intWinner {
							fmt.Printf("'%d','%d' MATCH!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n", intCard, intWinner)
							if cardScore == 0 {
								cardScore = 1
							} else {
							cardScore = cardScore*2
							}
						}
					}
				}
			}
		}
		// fmt.Printf("%s \n", winnerNumbers)
		// fmt.Printf("%s \n", cardNumbers)
		fmt.Printf("Card %d total score is %d\n", cardCount, cardScore)
		cardCount = cardCount + 1
		totalScore = totalScore + cardScore
		cardScore = 0
	}
	fmt.Printf("Total score: %d\n", totalScore)
}