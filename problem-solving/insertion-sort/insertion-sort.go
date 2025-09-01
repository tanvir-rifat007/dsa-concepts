package main

import "fmt"

func insertionSort(arr []int) []int {
	var pos, tmp int
	for i := 1; i < len(arr); i++ {
		tmp = arr[i]
		pos = i - 1

		for pos >= 0 {
			if tmp < arr[pos] {
				arr[pos+1] = arr[pos]
				pos--
			} else {
				break
			}

		}

		arr[pos+1] = tmp

	}

	return arr

}

func main() {
	arr := []int{100, -200, 0, 10, -20}

	fmt.Printf("%v\n", insertionSort(arr))

}
