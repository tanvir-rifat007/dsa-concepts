package main

import "fmt"

func binarySearch(arr []int, n int) int {
	left := 0
	right := len(arr) - 1
	var step int

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == n {

			fmt.Printf("step : %d\n", step)
			return mid
		} else if arr[mid] > n {

			right = mid - 1

			step++

		} else {
			left = mid + 1

			step++
		}

	}

	return -1
}

func main() {
	arr := []int{-5, -3, 0, 4, 20}
	fmt.Printf("The binary seach result is : %d\n", binarySearch(arr, -3))

}
