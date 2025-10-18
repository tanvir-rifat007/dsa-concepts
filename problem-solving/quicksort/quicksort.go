package main

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	var j int

	for j = low; j < high; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}

	}

	// after loop throw / when j < high then
	// swap the pivot with arr[i+1]
	swap(arr, i+1, j)

	return i + 1
}

func swap(arr []int, i, j int) {

	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func quicksort(arr []int, low, high int) {

	if low < high {

		pi := partition(arr, low, high)

		// pivots left side element

		quicksort(arr, low, pi-1)

		// pivots right side element
		quicksort(arr, pi+1, high)

	}

}

func main() {

	arr := []int{10, 7, 8, 9, 1, 5, 3, -100, 0}
	low := 0
	high := len(arr) - 1

	quicksort(arr, low, high)

	fmt.Printf("After sorting the array becomes : %v\n", arr)

}
