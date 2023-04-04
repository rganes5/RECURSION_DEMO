package main

import (
	"fmt"
)

func fib(num1 int, num2 int, limit int) {
	result := 0
	if limit > 0 {
		result = num1 + num2
		fmt.Printf(" %v", result)
		num1 = num2
		num2 = result
		fib(num1, num2, limit-1) //recursion
	}
}

func main() {
	var limit int
	num1 := 0
	num2 := 1
	fmt.Println("Enter the limit")
	fmt.Scan(&limit)
	fmt.Printf("The fibonocci series is as follows: ")
	fmt.Printf("%v %v", num1, num2)
	fib(num1, num2, limit-2)
}
