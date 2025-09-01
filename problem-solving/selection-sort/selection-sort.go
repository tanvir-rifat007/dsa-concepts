package main

import "fmt"

func selectionSort(arr []int) []int {

	var temp, lowestNumberIndex int

	for i := 0; i < len(arr)-1; i++ {
		lowestNumberIndex = i

		for j := i + 1; j < len(arr); j++ {

			if arr[j] < arr[lowestNumberIndex] {
				lowestNumberIndex = j // select lowestNumberIndeximum for each pass through(i)

			}

		}

		if lowestNumberIndex != i {

			temp = arr[lowestNumberIndex]
			arr[lowestNumberIndex] = arr[i]
			arr[i] = temp

		}

	}

	return arr

}

func main() {

	arr := []int{10, -29, -45, 0, 100}

	fmt.Printf("%v\n", selectionSort(arr))

}
