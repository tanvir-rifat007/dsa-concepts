package main

import "fmt"

func removeDuplicateFromArray(arr []int) []int {

	// create a hashmap
	exists := make(map[int]bool)

	for i, val := range arr {

		if exists[val] {
			// remove the duplicate element
			arr = append(arr[:i], arr[i+1:]...)

		}

		exists[val] = true

	}

	return arr
}

func main() {

	arr := []int{1, 10, 200, 2, 200}

	fmt.Printf("%d\n", removeDuplicateFromArray(arr))

}
