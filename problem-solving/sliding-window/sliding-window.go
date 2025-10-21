// [5,-2,3,1,2], B=3
// Pick first 3: 5,-2,3=6
// Slide window: pick first 2 and last 1: so sum: 6 - 3 + 2 = 5
// Slide window: pick first 1 and last 2: so sum: 5 - -2 + 1 = 8

package main

import "fmt"

func solve(A []int, B int) int {

	currSum := 0

	for i := 0; i < B; i++ {
		currSum += A[i]

	}

	maxSum := currSum

	for i := B - 1; i >= 0; i-- {
		currSum = currSum - A[i] + A[len(A)-B+i]
		if currSum > maxSum {
			maxSum = currSum
		}

	}

	return maxSum
}

func main() {

	fmt.Printf("%v\n", solve([]int{5, -2, 3, 1, 2}, 3))

}
