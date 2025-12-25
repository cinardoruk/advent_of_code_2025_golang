package main

import (
	// "bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	// "math"
	"os"
	// "strconv"
)

/*
* read each range from file
* get each id in the id range
* validate them
* add invalid ids

* write test for the example given
 */

// wrapper of convenience
func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Interval struct {
	start int
	end   int
}

func parseInput(filename string) []Interval {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	results := strings.Split(string(data), ",")

	var intervals []Interval
	var a, b int

	for _, value := range results {
		_, err := fmt.Sscanf(value, "%d-%d", &a, &b)
		if err != nil {
			// Handle error
		}
		intervals = append(intervals, Interval{a, b})

	}
	// log.Printf("\n%v\n", results)

	return intervals
}

// return false if id is invalid
func validateId(id string) bool {
	length := len(id)

	// anything repeated twice will be divisible by 2
	// conversely, if something isn't divisible by 2, it can't be a sequence repeated twice
	if length%2 != 0 {
		return true
	}

	// compare first half to second half
	if id[0:length/2] != id[length/2:] {
		return true
	}

	return false
}

// Now, an ID is invalid if it is made only of some sequence of digits repeated at least twice. So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times), and 1111111 (1 seven times) are all invalid IDs.
func validateId2(id string) bool {
	/*
		get id[0:0]
		if (the rest of the string just id[0] being repeated)
			return false
		get id[0:1]

	*/

	for i := 0; i < len(id)-1; i++ {
		// fmt.Printf("%d : %s\n", i, id[0:i+1])
		substring := id[0 : i+1]

		// if id is repeated copies of substring, len(id) is divisible by len(substring)
		if len(id)%len(substring) == 0 {
			// then, if id is substring*n, the below should be true
			if id == strings.Repeat(substring, len(id)/len(substring)) {
				return false
			}
		}

	}

	return true
}

func validateInterval(interval Interval) []int {
	var invalidIds []int
	for i := interval.start; i <= interval.end; i++ {
		if !validateId2(strconv.Itoa(i)) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func main() {
	// make slice of invalid ids
	var invalidIds []int
	// get all intervals from file
	intervals := parseInput("./input.txt")
	for _, interval := range intervals {
		invalidIds = append(invalidIds, validateInterval(interval)...)
	}

	sum := 0

	for _, id := range invalidIds {
		sum += id
	}

	fmt.Println(sum)
	fmt.Println(sum)
}
