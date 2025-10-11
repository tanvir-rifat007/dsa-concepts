// without stacking
// when cutting the chocolatebar

// time Complexity O(n*2)
// space is O(n) (because bars is growing with the input size)
package main

import "fmt"

func splitChocolateBar(n int) {
	bars := []int{n}

	step := 0

	for len(bars) > 0 {
		breakIndex := -1

		for i := 0; i < len(bars); i++ {
			if bars[i] > 1 {

				breakIndex = i // all time gets the 0 index
				break
			}

		}

		if breakIndex == -1 {
			break
		}

		size := bars[breakIndex]
		left := size / 2
		right := size - left

		fmt.Printf("Step %d: Break %d-piece bar → %d and %d\n", step, size, left, right)

		// empty the bars slice first time when bars = [8]
		// and on the second time bars = [4,4], so remove the first 4 and then add the
		// first 4's right(2) and left(2) on the second 4's
		// and get [4,2,2]
		bars = append(bars[:breakIndex], bars[breakIndex+1:]...)

		// then add the left and right to the bars slice
		bars = append(bars, left, right)

		// Show current state
		fmt.Printf("        Current pieces: %v\n", bars)

		step++

	}

	fmt.Printf("Total step : %d\n", step)

}

func main() {

	n := 8

	splitChocolateBar(n)

}
