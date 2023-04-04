package main

import (
	"fmt"
)

func sum(number int, tSum int) {
	if number > 0 {
		tSum += (number % 10)
		sum(number/10, tSum)
		fmt.Printf("\n%v\t%v", number, tSum)
	} else {
		fmt.Printf("\nThe sum of the digits is %v", tSum)
	}
}

func main() {
	var number int
	fmt.Println("Enter a number: ")
	fmt.Scan(&number)
	tSum := 0
	sum(number, tSum)

}
