package main

import "fmt"

// O(n^2)
// naive approach
func findDuplicate(arr []int) bool {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true

			}

		}

	}

	return false

}

// O(n)
// using hashmap

func findDuplicateUsingHashmap(arr []int) bool {
	hashmap := make(map[int]int)

	for i := 0; i < len(arr); i++ {

		hashmap[arr[i]] = (hashmap[arr[i]] + 1) | 0
	}

	for i := 0; i < len(arr); i++ {
		if hashmap[arr[i]] > 1 {
			return true

		}
	}

	return false
}

func main() {

	arr := []int{1, 2, 4, 2, 1, 10}

	fmt.Printf("%v\n", findDuplicateUsingHashmap(arr))

}
