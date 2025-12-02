package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

/*
	// read file
	// turn file lines into +/- ints
	// model the dial somehow
	// calculate how many times we hit 0
*/

// wrapper of convenience
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDialPosition(currentPosition int, input int) int {
	result := currentPosition + input

	if result > 0 {
		return result % 100
	} else {
		return (100 + result) % 100
	}
}

func getIntFromLine(line string) int {
	head := line[0]
	rest := line[1:]

	restNum, err := strconv.Atoi(rest)
	check(err)

	if string(head) == "L" {
		restNum = -restNum
	}

	return restNum
}

func testFunctions() {
	initialDialPosition := 50
	currentDialPosition := initialDialPosition

	turns := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	correctPositions := []int{
		82,
		52,
		0,
		95,
		55,
		0,
		99,
		0,
		14,
		32,
	}

	for i := range turns{
		// fmt.Println(turns[i])
		motion := getIntFromLine(turns[i])

		currentDialPosition = getDialPosition(currentDialPosition, motion)
		fmt.Printf("move: %d reaches: %d", motion, currentDialPosition)
		fmt.Printf(" correct : %d\n", correctPositions[i] )
	}
}

func processInput(){
	initialDialPosition := 50
	currentDialPosition := initialDialPosition

	numZeroes := 0

	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	// get a 'scanner'
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()

		currentDialPosition = getDialPosition(currentDialPosition, getIntFromLine(line))
		// fmt.Println(currentDialPosition)
		if currentDialPosition == 0 {
			numZeroes++
		}
	}

	// Check for errors during scanning
	check(scanner.Err())

	fmt.Println(numZeroes)
}

func main() {
	//somehow need this for neovim to show stdout correctly
	fmt.Println()

	// testFunctions()

	processInput()
}

/*
how am i supposed to model the dial?
is it just mod?
97 98 99 0 1 2 3 4 5 6
   ^
result = cur + motion
if result > 0
	return result % 100
else
	return (100 + result) % 100

*/
