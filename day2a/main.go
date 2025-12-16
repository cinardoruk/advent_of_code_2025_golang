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

func validateInterval(interval Interval) []int {
	var invalidIds []int
	for i := interval.start; i <= interval.end; i++ {
		if !validateId(strconv.Itoa(i)) {
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
