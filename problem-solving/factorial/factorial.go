package main

import "fmt"

// TCO(tail call optimization)

func factorial(n int, acc int) int {

	if n == 1 {
		return acc
	}
	return factorial(n-1, n*acc)
}

func main() {

	res := factorial(5, 1)

	fmt.Printf("The result : %d\n", res)

}
