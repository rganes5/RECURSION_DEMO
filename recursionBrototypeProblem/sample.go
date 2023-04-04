package main

import "fmt"

func main() {
	n := 5
	functionSample(n)
}

func functionSample(n int) {
	if n <= 1 {
		return
	}
	functionSample(n - 1)
	fmt.Printf("%v\t", n)
	functionSample(n - 1)
}
