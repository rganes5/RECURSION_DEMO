package main

import (
	"fmt"
)

func fact(number int) int {
	if number <= 1 {
		return 1
	}
	return number * fact(number-1) // RECURSION
}

func main() {
	var number int
	fmt.Println("Enter a number: ")
	fmt.Scan(&number)
	x := fact(number)
	fmt.Printf("%v", x)
}
