package main

import "fmt"

func binarySearch(arr []int, n int) int {
	first := 0
	last := len(arr) - 1

	for first <= last {
		mid := (first + last) / 2
		if arr[mid] == n {
			return mid
		} else if arr[mid] < n {
			first = mid + 1

		} else {
			last = mid - 1
		}

	}

	return -1
}

func main() {

	arrayOfNumbers := []int{-5, -3, 0, 4, 20}
	fmt.Printf("%d\n", binarySearch(arrayOfNumbers, -4))

}
