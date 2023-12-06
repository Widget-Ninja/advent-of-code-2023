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
	input, err := os.Open("C:/Users/shamu/Documents/Code/advent-of-code-2023/day4/inputTest.txt")
	check(err)

	fileScan := bufio.NewScanner(input)
    fileScan.Split(bufio.ScanLines)
	var inputLines []string
    for fileScan.Scan() {
        inputLines = append(inputLines, fileScan.Text())
    }
	var totalScore int
	var totalCount []int
	cardCount := 1
	for _, line := range inputLines {
		var lineSplit []string
		lineSplit = strings.Split(line, ": ")
		lineSplit = strings.Split(lineSplit[1], " | ")
		winner := lineSplit[0]
		winnerNumbers := strings.Split(winner, " ")
		card := lineSplit[1]
		cardNumbers := strings.Split(card, " ")
		cardScore := 0
		for _, cardNumber := range cardNumbers {
			var intCard int
			if cardNumber != "" {
				// fmt.Printf("card number '%s'\n", cardNumber)
				intCard, err = strconv.Atoi(cardNumber)
				check(err)
			}
			for _, winnerNumber := range winnerNumbers {
				var intWinner int
				if winnerNumber != "" {
					intWinner, err = strconv.Atoi(winnerNumber)
					check(err)
				}
				if intCard != 0 {
					if intWinner != 0 {
						if intCard == intWinner {
							fmt.Printf("'%d','%d' MATCH!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n", intCard, intWinner)
								cardScore = cardScore + 1
						}
					}
				}
			}
		}
		fmt.Printf("totalCount is: %d \n", totalCount)
		var totalCountAmmend []int
		if len(totalCount) < cardCount + cardScore {
			var appendLoop int 
			var arrayAppend []int
			appendLoop = cardCount + cardScore - len(totalCount)
			for i := 1; i < appendLoop; i++ {
				arrayAppend[0] = 0
				totalCountAmmend = append(totalCountAmmend, arrayAppend...)
			}
			totalCount = append(totalCount, totalCountAmmend...)
		}
		fmt.Printf("totalCount is: %d \n", totalCount)
		fmt.Printf("Card %d total score is %d\n", cardCount, cardScore)
		for currentCount := 1; currentCount < cardScore + 1; currentCount++ {
			totalCount[cardCount + currentCount] = totalCount[cardCount + currentCount] + 1
		}
		cardCount = cardCount + 1
		// cardScore = 0
	}
	for _, counts := range totalCount {
		totalScore = totalScore + counts
	}
	fmt.Printf("Total score: %d\n", totalScore)
}