package main

import (
	"fmt"
	"os"
	"regexp"
	"bufio"
	"strings"
	"strconv"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func convertNumberText(str string) string {
	m0 := regexp.MustCompile(`zero`)
	m1 := regexp.MustCompile(`one`)
	m2 := regexp.MustCompile(`two`)
	m3 := regexp.MustCompile(`three`)
	m4 := regexp.MustCompile(`four`)
	m5 := regexp.MustCompile(`five`)
	m6 := regexp.MustCompile(`six`)
	m7 := regexp.MustCompile(`seven`)
	m8 := regexp.MustCompile(`eight`)
	m9 := regexp.MustCompile(`nine`)
	str = m0.ReplaceAllString(str, "0")
	str = m1.ReplaceAllString(str, "1")
	str = m2.ReplaceAllString(str, "2")
	str = m3.ReplaceAllString(str, "3")
	str = m4.ReplaceAllString(str, "4")
	str = m5.ReplaceAllString(str, "5")
	str = m6.ReplaceAllString(str, "6")
	str = m7.ReplaceAllString(str, "7")
	str = m8.ReplaceAllString(str, "8")
	str = m9.ReplaceAllString(str, "9")
	return str
}

func main () {
	input, err := os.Open("C:/Users/shamu/Documents/Code/advent-of-code-2023/day1/input.txt")
	check(err)
	re := regexp.MustCompile("[0-9]+")
    fileScan := bufio.NewScanner(input)
    fileScan.Split(bufio.ScanLines)
    var coordinatesLines []string
	var numbers string
	var finalCoordinates int

    for fileScan.Scan() {
        coordinatesLines = append(coordinatesLines, fileScan.Text())
    }

	for _, value := range coordinatesLines {
		var intConversion int
		fmt.Printf("before text conversion: %s\n",value)
		numbers = convertNumberText(value)
		// fmt.Printf("after text conversion: %s\n",numbers)
		numbers = strings.Join(re.FindAllString(string(numbers), -1), numbers)
		// fmt.Printf("pull numbers: %s\n",numbers)
		numbers = string(numbers[0]) + string(numbers[len(numbers)-1])
		fmt.Printf("combine first and last: %s\n",numbers)
		intConversion, err = strconv.Atoi(string(numbers))
		check(err)
		finalCoordinates = finalCoordinates + intConversion
	}
	fmt.Print(finalCoordinates)
}