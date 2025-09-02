package main

import "fmt"

/*
func maxOfArray(arr []int) int {
	fmt.Println("Recursion")
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	var res = maxOfArray(arr[1:])

	if arr[0] > res {
		return arr[0]
	}

	return res

}
*/

func fibo(n int, hash map[int]int) int {

	if n == 0 || n == 1 {
		return n
	}

	if hash[n] == 0 {
		hash[n] = fibo(n-1, hash) + fibo(n-2, hash)

	}

	return hash[n]

}

func main() {
	fmt.Printf("%v\n", fibo(10, make(map[int]int)))

}
