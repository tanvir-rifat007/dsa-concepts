package main

import "fmt"

func mergeSort(arr []int) []int {

	// base case:

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	fmt.Println(left)
	fmt.Println(right)

	return merge(left, right)

}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++

		} else {
			result = append(result, right[j])

			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}

func main() {

	fmt.Printf("%v\n", mergeSort([]int{0, -2, 4, 1}))
}
