package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	test := "abcd"
	for i := range test{
		fmt.Printf("%d : %s\n", i, test[0:i+1])
	}
}
