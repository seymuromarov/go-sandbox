package main

import "fmt"

func main() {
	number := 10
	n1, n2 := 0, 1
	var nextTerm int

	fmt.Println("Fibonacci Series in Go:")

	for i := 1; i <= number; i++ {
		fmt.Println(n1)
		nextTerm = n1 + n2
		n1 = n2
		n2 = nextTerm
	}
}
