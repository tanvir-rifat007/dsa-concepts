package main

import "fmt"

// exponentiation time complexity: O(2^n)
// space complexity O(n)

// func fibo(n int) int {

// 	if n <= 1 {
// 		return n

// 	}
// 	return fibo(n-1) + fibo(n-2)

// }

// time complexity O(n)
// space complexity O(1)
func fibo(n int) int {
	a, b := 0, 1

	var temp int
	for i := 2; i <= n; i++ {
		temp = a + b
		a = b
		b = temp

	}

	return temp

}

func main() {

	res := fibo(6)
	fmt.Printf("The fib of 5 is : %d\n", res)

}
